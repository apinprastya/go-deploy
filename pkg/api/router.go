package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) checkSecret(ctx *gin.Context) {
	secret := ctx.GetHeader("SECRET")
	if secret != a.secret {
		ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"error": "you have no access"})
	}
}

func (a *Api) RegisterRouter() {

	a.ginEngine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "page not found"})
	})

	a.ginEngine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "method not allowed"})
	})

	a.ginEngine.POST("/upload", a.checkSecret, func(ctx *gin.Context) {
		version := ctx.PostForm("version")
		path := ctx.PostForm("path")
		filename := ctx.PostForm("filename")
		ff, err := ctx.FormFile("file")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		f, err := ff.Open()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		a.srv.SaveFile(version, path, filename, f)
		ctx.JSON(http.StatusCreated, map[string]interface{}{"status": true})
	})

	a.ginEngine.GET("/setlive/:version", a.checkSecret, func(ctx *gin.Context) {
		version := ctx.Param("version")
		if err := a.srv.SetLive(version); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, map[string]interface{}{"status": true})
	})
}
