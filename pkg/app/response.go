package app

import (
	"context"
	"encoding/json"
	"net/http"
	"onecho_sso_backend/pkg/enums"
	"onecho_sso_backend/pkg/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Gin struct {
	C              *gin.Context
	RequestContext context.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (res *Response) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return nil
}

func (res Response) asString() string {
	strByte, _ := json.Marshal(res)
	return string(strByte)
}

// Response setting gin.JSON
func (g *Gin) Response(errCode int, data interface{}) {
	result := Response{
		Code: errCode,
		Msg:  enums.GetMsg(errCode),
		Data: data,
	}
	g.C.JSON(http.StatusOK, result)

	// record response
	logging.Ctx(g.RequestContext).Debug("", zap.String("response", result.asString()))

	return
}
