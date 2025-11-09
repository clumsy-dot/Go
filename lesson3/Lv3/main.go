package main

import (
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name  string    `json:"name"`
	Score []float64 `json:"score"`
}

func main() {
	r := gin.Default()
	r.POST("/average", Average)
	r.Run(":80")
}

func Average(r *gin.Context) {
	var s Student
	var average float64
	var sum float64

	err := r.ShouldBindJSON(&s)

	if err != nil {
		r.JSON(400, gin.H{
			"error": "错误",
		})
	}
	for _, v := range s.Score {
		sum += v
	}

	average = sum / float64(len(s.Score))

	r.JSON(200, gin.H{
		"average": average,
	})
}
