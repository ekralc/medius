package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(VerifyRequestMiddleware())
	r.POST("/interactions", interactionHandler)

	r.Run()
}
