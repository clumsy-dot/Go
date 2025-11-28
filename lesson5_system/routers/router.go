package routers

import (
	"system/api"
	"system/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()

	public := r.Group("/选课平台")
	{
		public.POST("/register", api.Register)
		public.POST("/login", api.Login)
		public.POST("/refresh-token", api.RefreshToken)
		public.POST("/course", api.GetCourse)
	}
	auth := r.Group("/选课平台")
	auth.Use(middleware.Auth())
	{
		auth.POST("/enroll", api.Enrollment)
		auth.DELETE("/enroll", api.DropCourse)
		auth.GET("/enrollments", api.GetUserEnrollments)
		admin := auth.Group("/admin")
		admin.Use(middleware.AuthOnly())
		{
			admin.POST("/creatcourse", api.CreateCourse)
		}
	}
	return r
}
