package api

import (
	"fmt"
	"godeploy/pkg/service"

	"github.com/gin-gonic/gin"
)

type Api struct {
	srv       *service.Service
	ginEngine *gin.Engine
	secret    string
}

func NewApi(secret string, srv *service.Service) *Api {
	return &Api{secret: secret, srv: srv, ginEngine: gin.Default()}
}

func (a *Api) Run(address string, port int) error {
	fmt.Println("Running API")

	a.RegisterRouter()

	a.ginEngine.Run(fmt.Sprintf("%s:%d", address, port))

	return nil
}
