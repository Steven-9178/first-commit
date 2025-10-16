package main

import (
	"github.com/gin-gonic/gin"
	"web/server"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", server.Register)

	engine.Run(":8888")
}
