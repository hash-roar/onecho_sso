package main

import (
	"fmt"
	"net/http"
	"onecho_sso_backend/models"
	"onecho_sso_backend/pkg/gredis"
	"onecho_sso_backend/pkg/logging"
	"onecho_sso_backend/pkg/setting"
	"onecho_sso_backend/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	// logging.Setup()
	models.Setup()
	gredis.Setup()
}

func main() {
	// gin.SetMode()
	gin.SetMode(setting.ServerSetting.RunMode)
	router := routers.InitRouter()

	listen := fmt.Sprintf("%s", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:           listen,
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	logging.Infof("server listen at [%s]", listen)

	server.ListenAndServe()
}
