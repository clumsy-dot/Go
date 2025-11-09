package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.StaticFile("/cat.jpg", "./cat.jpg")
	//c.StaticFile("/b.jpg", "./b.jpg")
	//c.StaticFile("/a.jpg", "./a.jpg")
	c.Run(":80")
}
