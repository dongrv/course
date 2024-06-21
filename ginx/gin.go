package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	r.Use(authMiddleware, returnZapLoggerMiddleware(logger))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8086"); err != nil {
		fmt.Printf("[GIN]Server run error found:%s", err.Error())
		return
	}
}

func authMiddleware(c *gin.Context) {
	username, ok := c.GetQuery("username")
	if !ok || username != "Tony" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "username needed",
		})
		return
	}
	println("当前用户：", username)
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
