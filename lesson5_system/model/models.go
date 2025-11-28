/*所有层要用的结构体*/

package model

import (
	"time"
)

// 用户  区分管理员和普通学生
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"` //主键
	Username  string    `gorm:"size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Role      string    `gorm:"size:20;default:student" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 课程
type Course struct {
	ID         uint      `gorm:"primaryKey" json:"id"` //主键
	Name       string    `gorm:"size:100;not null" json:"name"`
	Capacity   int       `gorm:"not null" json:"capacity"`
	CurrentNum int       `gorm:"defalut:0" json:"current_num"`
	Teacher    string    `gorm:"size:50" json:"teacher"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// 抢课
type Enrollment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`            //抢课记录ID  主键
	UserID    uint      `gorm:"not null;index" json:"user_id"`   //学生ID
	CourseID  uint      `gorm:"not null;index" json:"course_id"` //课程ID
	CreatedAt time.Time `json:"created_at"`

	User   User   `gorm:"foreignKey:UserID" json:"user"`
	Course Course `gorm:"foreignKey:CourseID" json:"course"`
}

// 黑名单
type TokenBlacklist struct {
	ID    uint      `gorm:"primaryKey" json:"id"` //退出记录  主键
	Token string    `gorm:"size:50;uniqueIndex" json:"Token"`
	ExpAt time.Time `json:"exp_at"`
}
