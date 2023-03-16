package api

import (
	"onecho_sso_backend/pkg/app"
	"onecho_sso_backend/pkg/enums"
	"onecho_sso_backend/pkg/logging"
	emailservice "onecho_sso_backend/services/email_service"
	userservice "onecho_sso_backend/services/user_service"
	"onecho_sso_backend/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type newUserRegisterForm struct {
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	NickName     string `json:"nick_name"`
	Password     string `json:"password" binding:"required"`
	ValidateCode string `json:"validate_code"`
}

func UserRegister(c *gin.Context) {
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
	var code int
	if form.Email != "" && form.Password != "" {
		// validate email
		if emailservice.ValidateEmail(App.RequestContext, form.Email, form.ValidateCode) {
			code = userservice.RegisterByEmail(App.RequestContext, form.Email, form.Password)
		} else {
			code = enums.EMAIL_VALIDATION_ERROR
		}
	} else if form.NickName != "" && form.Password != "" {
		code = userservice.RegisterByName(App.RequestContext, form.NickName, form.Password)
	} else {
		code = enums.INVALID_PARAMS
	}
	App.Response(code, "")
}
