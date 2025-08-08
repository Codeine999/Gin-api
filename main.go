package main

import (
	"gin/initializers"
	"gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.ConnectDb()

	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	routers.Register(router)
	router.Run(":666")
}
