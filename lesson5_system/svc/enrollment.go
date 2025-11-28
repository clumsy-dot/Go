package svc

import (
	"fmt"
	"system/dao"
	"system/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 抢课
func EnrollCourse(db *gorm.DB, UserID uint, CourseID uint) error {
	tx := dao.DB.Begin()
	if tx.Error != nil {
		fmt.Println("开启事务失败", tx.Error)
		return tx.Error
	}
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()
	var course model.Course
	err1 := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id=?", CourseID).First(&course).Error
	if err1 != nil {
		tx.Rollback()
		if err1 == gorm.ErrRecordNotFound {
			return fmt.Errorf("未找到课程:%v", CourseID)
		}
		return fmt.Errorf("查询课程失败: %v", err1)
	}
	if course.CurrentNum >= course.Capacity {
		tx.Rollback()
		return fmt.Errorf("课程【%s】已满 (容量:%d,当前人数:%d)",
			course.Name, course.Capacity, course.CurrentNum)
	}
	err2 := tx.Model(&model.Course{}).Where("id=?", CourseID).Update("current_num", course.CurrentNum+1).Error
	if err2 != nil {
		tx.Rollback()
		return fmt.Errorf("更新选课人数失败:%v", err2)
	}
	sc := model.Enrollment{
		UserID:   UserID,
		CourseID: CourseID,
	}
	err3 := tx.Create(&sc).Error
	if err3 != nil {
		tx.Rollback()
		return fmt.Errorf("选课失败: %v", err3)
	}
	err4 := tx.Commit().Error
	if err4 != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %v", err4)
	}

	return nil
}

// 退课
func DropCourse(userID, courseID uint) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("user_id = ? AND course_id = ?", userID, courseID).Delete(&model.Enrollment{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("未找到选课记录")
		}

		var course model.Course
		if err := tx.First(&course, courseID).Error; err != nil {
			return err
		}

		if course.CurrentNum > 0 {
			if err := tx.Model(&course).Update("current_num", course.CurrentNum-1).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// 获取抢课课表
func GetUserEnrollments(userID uint) ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	result := dao.DB.Where("user_id = ?", userID).Find(&enrollments)
	if result.Error != nil {
		return nil, result.Error
	}
	return enrollments, nil
}
