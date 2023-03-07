package api

import (
	"onecho_sso_backend/pkg/app"
	"onecho_sso_backend/pkg/enums"
	"onecho_sso_backend/pkg/logging"
	"onecho_sso_backend/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type newUserRegisterForm struct {
	Email        string
	Phone        string
	Password     string
	ValidateCode string
}

func NewUserRegister(c *gin.Context) {
	App := app.Gin{
		C:              c,
		RequestContext: utils.GenerateRequestContext(c),
	}
	form := new(newUserRegisterForm)
	if err := c.ShouldBindJSON(form); err != nil {
		logging.Ctx(App.RequestContext).Info("get param error", zap.Error(err))
		App.Response(enums.INVALID_PARAMS, "")
		return
	}
	App.Response(enums.SUCCESS, "")
}
