package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/server"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", server.Register)

	fmt.Println("我是hyx")

	engine.Run(":8888")
}
