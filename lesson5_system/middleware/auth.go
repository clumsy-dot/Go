package middleware

import (
	"fmt"
	"system/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHadler := c.GetHeader("Authorization")
		if authHadler == "" {
			utils.Error(c, 401, "请登录")
			c.Abort()
			return
		}
		claim, err := utils.VerifyAccessToken(authHadler)
		fmt.Println(claim)
		if err != nil {
			utils.Error(c, 403, "Token错误")
			c.Abort()
			return
		}
		c.Set("user_id", claim.UserID)
		c.Set("user_role", claim.Role)
		c.Next()
	}
}

func AuthOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role.(string) != "admin" {
			utils.Error(c, 401, "你不是管理员")
			c.Abort()
			return
		}
		c.Next()
	}
}
