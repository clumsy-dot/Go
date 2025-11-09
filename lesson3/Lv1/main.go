package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/talk", ping)
	r.Run(":80")
}

func ping(a *gin.Context) {
	msg := a.Query("msg")
	var date string
	switch msg {
	case "ping":
		date = "pong"
	case "helloserver":
		date = "helloclient"
	default:
		date = "error"
	}
	a.JSON(http.StatusOK, gin.H{
		"date": date,
	})
}
