package api

import (
	"onecho_sso_backend/pkg/app"
	"onecho_sso_backend/pkg/enums"
	"onecho_sso_backend/pkg/logging"
	emailservice "onecho_sso_backend/services/email_service"
	"onecho_sso_backend/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type registerValidationForm struct {
	Email string `binding:"required,email"`
}

func GetRegisterValidationEmail(c *gin.Context) {
	App := app.Gin{
		C:              c,
		RequestContext: utils.GenerateRequestContext(c),
	}
	form := new(registerValidationForm)
	if err := c.ShouldBindJSON(form); err != nil {
		logging.Ctx(App.RequestContext).Info("get param error", zap.Error(err))
		App.Response(enums.INVALID_PARAMS, "")
		return
	}
	var code int
	code = emailservice.SendValidationEmail(App.RequestContext, form.Email)
	App.Response(code, "")
}
