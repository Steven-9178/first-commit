package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/server"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", server.Register) /////cnm

	fmt.Println("我是hyx")
	fmt.Println("我是sb")

	engine.Run(":8888")
}
