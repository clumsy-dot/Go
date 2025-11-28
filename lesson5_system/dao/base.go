/*初始化连接*/
package dao

import (
	"fmt"
	"system/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:20070610li@tcp(127.0.0.1:3306)/学生选课数据库?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接错误", err)
		return
	}
	db.AutoMigrate(&model.Course{}, &model.User{}, &model.Enrollment{})

	DB = db
}
