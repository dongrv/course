package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"time"
)

// 获取日志管理器
func initZapLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()                         // NewProductionConfig
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder // 自定义时间编码器
	config.Level.SetLevel(zapcore.InfoLevel)

	// 初始化 Zap logger
	logger, err := config.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func Run() {
	logger, err := initZapLogger()
	if err != nil {
		println("Init logger err:", err.Error())
		return
	}
	defer logger.Sync() // 程序退出时确保日志刷新到存储介质

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.Use(
		authMiddleware,
		returnZapLoggerMiddleware(logger),
		securityHeaderMiddleware,
	)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	type LoginForm struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	r.POST("/bind_form", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		if c.ShouldBind(&form) != nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": "unahthorized"})
			}
		}
	})

	r.POST("/post_form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "psoted",
			"message": message,
			"nick":    nick,
		})
	})

	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>", // 提供字面字符
		})
	})

	r.POST("/post", func(c *gin.Context) {
		//POST /post?id=1234&page=1 HTTP/1.1
		//Content-Type: application/x-www-form-urlencoded
		//
		//name=manu&message=this_is_great
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id:%s; page:%s; name:%s; message:%s\n", id, page, name, message)
	})

	// 使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names) // 将输出：while(1);["lena","austin","foo"]
	})

	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// 使用多种格式输出结果
	r.GET("/jsonResult", func(c *gin.Context) {
		// gin.H 是 map[string]interface{} 的一种快捷方式
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "hey"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"status": http.StatusOK, "message": "hey"})
	})
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": "hey"})
	})
	r.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{1, 2}
		label := "test"

		type ProtoStruct struct {
			proto.Message
			Label string
			Reps  []int64
		}

		data := &ProtoStruct{
			Label: label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 ProtoStruct protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data) // 会报错，需要真实的protobuf结构体
	})

	// 绑定uri
	type Person struct {
		ID   string `uri:"id" binding:"required,uuid"`
		Name string `uri:"name" binding:"required"`
	}

	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSONP(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})

	// 多次绑定
	r.GET("/multibind", func(c *gin.Context) {
		type formA struct {
			Foo string `json:"foo" xml:"foo" binding:"required"`
		}
		type formB struct {
			Bar string `json:"bar" xml:"bar" binding:"required"`
		}
		a := &formA{}
		b := &formB{}
		// ShouldBindBodyWith 可以复用c.Request.Body
		if errA := c.ShouldBindBodyWith(a, binding.JSON); errA == nil {
			c.String(http.StatusOK, "the body should be formA")
		} else if errB := c.ShouldBindBodyWith(b, binding.JSON); errB == nil {
			c.String(http.StatusOK, "the body should be JSON formB")
		} else if errB := c.ShouldBindBodyWith(b, binding.XML); errB == nil {
			c.String(http.StatusOK, "the body should be XML formB")
		}
	})

	if err := r.Run(":8086"); err != nil {
		fmt.Printf("[GIN]Server run error found:%s", err.Error())
		return
	}
}

func authMiddleware(c *gin.Context) {
	/*username, ok := c.GetQuery("username")
	if !ok || username != "Tony" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "username needed",
		})
		return
	}
	println("当前用户：", username)*/
	c.Next()
}

func returnZapLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		logger.Info("HTTP request",
			zap.Int("status", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("ip", clientIP),
			zap.String("latency", latency.String()),
		)
	}
}

var expectedHost = "localhost:8086"

func securityHeaderMiddleware(c *gin.Context) {
	if c.Request.Host != expectedHost {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
		return
	}
	c.Header("X-Frame-Options", "DENY")
	c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	c.Header("Referrer-Policy", "strict-origin")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	c.Next()
}
