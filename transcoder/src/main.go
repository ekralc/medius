package main

import (
	"bufio"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/pubsub", PubsubPush)
	r.Run()
}

func handleReader(reader *bufio.Reader) {
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Print(str)
	}
}
