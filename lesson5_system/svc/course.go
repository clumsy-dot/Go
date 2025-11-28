package svc

import (
	"system/dao"
	"system/model"
)

type CourseRep struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" binding:"required"`
	Teacher  string `json:"teacher"`
	Capacity int    `json:"capacity" binding:"required"`
}

// 创建课程
func CreateCourse(rep *CourseRep) (*model.Course, error) {
	course := &model.Course{
		ID:         rep.ID,
		Name:       rep.Name,
		Teacher:    rep.Teacher,
		Capacity:   rep.Capacity,
		CurrentNum: 0,
	}
	res := dao.DB.Create(course)
	if res.Error != nil {
		return nil, res.Error
	}
	return course, nil
}

// 获取课程表
func GetCourse() ([]model.Course, error) {
	var courses []model.Course
	result := dao.DB.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}
