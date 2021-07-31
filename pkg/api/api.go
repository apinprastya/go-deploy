package api

import (
	"fmt"
	"godeploy/pkg/service"

	"github.com/gin-gonic/gin"
)

type Api struct {
	srv    *service.Service
	r      *gin.Engine
	secret string
}

func NewApi(secret string, srv *service.Service) *Api {
	return &Api{secret: secret, srv: srv, r: gin.Default()}
}

func (a *Api) Run(address string, port int) error {
	fmt.Println("Running API")

	a.RegisterRouter()

	a.r.Run(fmt.Sprintf("%s:%d", address, port))

	return nil
}
