package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.Handle("GET", "/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Println("engine error", engine.Run(":8080"))
}
