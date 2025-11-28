package api

import (
	"system/svc"
	"system/utils"

	"github.com/gin-gonic/gin"
)

// 创建用户
func Register(c *gin.Context) {
	var rep svc.RepUser
	err := c.ShouldBindJSON(&rep)
	if err != nil {
		utils.Error(c, 400, "参数错误")
		return
	}
	user, err := svc.CreateUser(&rep)
	if err != nil {
		utils.Error(c, 500, err.Error())
	}
	utils.Success(c, user)
}

// 登录
func Login(c *gin.Context) {
	var req svc.RepLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	Atoken, Rtoken, User, err := svc.Login(&req)
	if err != nil {
		utils.Error(c, 401, "用户名或密码错误")
		return
	}

	utils.Success(c, gin.H{
		"accesstoken":  Atoken,
		"refreshToken": Rtoken,
		"user":         User.Username,
	})
}

func RefreshToken(c *gin.Context) {
	var rep struct {
		Token string `json:"refreshtoken"`
	}
	err := c.ShouldBindJSON(&rep)
	if err != nil {
		utils.Error(c, 400, "错误")
		return
	}
	NewAToken, NewRToken, err := utils.RefreshToken(rep.Token)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}
	utils.Success(c, gin.H{
		"accessToken":  NewAToken,
		"refreshToken": NewRToken,
	})
}
