package routers

import (
	middleware "onecho_sso_backend/midlleware"
	"onecho_sso_backend/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.DefaultLogger())

	authGroup := r.Group("auth")
	{
		authGroup.POST("register", api.UserRegister)
		authGroup.POST("reg-email", api.GetRegisterValidationEmail)
	}
	return r
}
