/*回应*/
package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
