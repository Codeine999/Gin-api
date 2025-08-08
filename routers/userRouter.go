package routers

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Wake up Codeine",
		})
	})
}
