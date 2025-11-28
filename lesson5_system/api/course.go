package api

import (
	"system/svc"
	"system/utils"

	"github.com/gin-gonic/gin"
)

// 创建课程
func CreateCourse(c *gin.Context) {
	var rep svc.CourseRep
	err := c.ShouldBindJSON(&rep)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	course, err := svc.CreateCourse(&rep)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	utils.Success(c, course)
}

// 获取课程
func GetCourse(c *gin.Context) {
	course, err := svc.GetCourse()
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	utils.Success(c, course)
}
