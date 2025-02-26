package main

import (
	"com/chat/service/cmd/chat/initial"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	initial.InitApp()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8077")
	if err != nil {
		fmt.Println("1231312")
		fmt.Println(err)
		os.Exit(1)
	}
}
