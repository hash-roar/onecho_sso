package utils

import (
	"context"
	"onecho_sso_backend/pkg/enums"

	"github.com/gin-gonic/gin"
)

type RequestContext struct {
	RequestId string
}

func GenerateRequestContext(c *gin.Context) context.Context {
	return context.WithValue(c.Request.Context(), "reqctx", RequestContext{
		RequestId: c.MustGet(enums.RequestIdName).(string),
	})
}
