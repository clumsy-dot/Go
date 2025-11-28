package svc

import (
	"system/dao"
	"system/model"
	"system/utils"
)

type RepUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type RepLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 创建用户
func CreateUser(rep *RepUser) (*model.User, error) {
	user := &model.User{
		ID:       rep.ID,
		Username: rep.Username,
		Password: rep.Password,
		Role:     rep.Role,
	}
	res := dao.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// 登录
func Login(rep *RepLogin) (string, string, *model.User, error) {
	var user model.User
	res := dao.DB.Where("username = ? AND password = ?", rep.Username, rep.Password).First(&user)
	if res.Error != nil {
		return "", "", nil, res.Error
	}
	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, user.Role)
	if err != nil {
		return "", "", nil, err
	}
	return accessToken, refreshToken, &user, nil
}
