package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Course struct {
	gorm.Model
	Name       string
	Capacity   int
	CurrentNum int
}

type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade int
}

type StudentCourse struct {
	gorm.Model
	StudentID uint `gorm:"uniqueIndex:idx_stu_course"`
	CourseID  uint `gorm:"uniqueIndex:idx_stu_course"`
}

var db *gorm.DB

func main() {
	gin.SetMode("release")
	r := gin.Default()
	dsn := "root:20070610li@tcp(127.0.0.1:3306)/学生选课数据库?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接错误", err)
		return
	}

	if mErr := db.AutoMigrate(&Student{}, &Course{}, &StudentCourse{}); mErr != nil {
		fmt.Println("自动迁移失败", mErr)
		return
	}
	r.POST("/选课平台/新建学生", f1)
	r.POST("/选课平台/新建课程", f2)
	r.POST("/选课平台", f3)
	r.Run(":8080")
}

func CreateStudent(db *gorm.DB, name string, age int, garde int) error {
	student := Student{Name: name, Age: age, Grade: garde}
	res := db.Create(&student)
	if res.Error != nil {
		fmt.Println("新建学生失败", res.Error)
		return res.Error
	}
	return nil
}

func CreateCourse(db *gorm.DB, name string, capacity int) error {
	course := Course{
		Name:       name,
		Capacity:   capacity,
		CurrentNum: 0}
	res := db.Create(&course)
	if res.Error != nil {
		fmt.Println("新建课程失败", res.Error)
		return res.Error
	}
	return nil
}

func EnrollCourse(db *gorm.DB, StudentID uint, CourseID uint) error {
	tx := db.Begin()
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
	var course Course
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
	err2 := tx.Model(&Course{}).Where("id=?", CourseID).Update("current_num", course.CurrentNum+1).Error
	if err2 != nil {
		tx.Rollback()
		return fmt.Errorf("更新选课人数失败:%v", err2)
	}
	sc := StudentCourse{
		StudentID: StudentID,
		CourseID:  CourseID,
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

func f1(c *gin.Context) {
	var student Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err2 := CreateStudent(db, student.Name, student.Age, student.Grade)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

func f2(c *gin.Context) {
	var course Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err2 := CreateCourse(db, course.Name, course.Capacity)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err2.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

func f3(c *gin.Context) {
	var ID StudentCourse
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err3 := EnrollCourse(db, ID.StudentID, ID.CourseID)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err3.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "抢课成功"})
}
