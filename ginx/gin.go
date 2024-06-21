package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(authMiddleware)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8086")
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
