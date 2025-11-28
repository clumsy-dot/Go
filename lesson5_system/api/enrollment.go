package api

import (
	"strconv"
	"system/dao"
	"system/svc"
	"system/utils"

	"github.com/gin-gonic/gin"
)

// 抢课
func Enrollment(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	courseID, _ := strconv.Atoi(c.Query("course_id"))
	err := svc.EnrollCourse(dao.DB, uint(userID), uint(courseID))
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, "选课成功")
}

// 退课
func DropCourse(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	courseID, _ := strconv.Atoi(c.Query("course_id"))
	err := svc.DropCourse(uint(userID), uint(courseID))
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, "退课成功")
}

// 获取抢课课表
func GetUserEnrollments(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	enrollments, err := svc.GetUserEnrollments(uint(userID))
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, enrollments)
}
