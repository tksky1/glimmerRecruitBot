package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Config struct {
	recruitGroupId int64
	adminGroupId   int64
	listenAt       int
	postTo         int
}

var config Config = Config{
	recruitGroupId: 683234808,
	adminGroupId:   202577501,
	listenAt:       5701,
	postTo:         5700,
}

func main() {
	r := gin.Default()
	r.POST("/", handlePost)
	println("-- Glimmer Recruit Bot --")
	println("Listen at", config.listenAt, "; Post to", config.postTo)
	err := r.Run(":" + strconv.Itoa(config.listenAt))
	if err != nil {
		println(err.Error())
	}
}
