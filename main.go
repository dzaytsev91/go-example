package main

import (
	"github.com/dzaytsev91/go-example/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run()
}
