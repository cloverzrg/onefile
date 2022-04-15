package api

import (
	"github.com/cloverzrg/onefile/api/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.GET("/login", controller.Login)
	r.GET("/callback", controller.Callback)
	r.GET("/token", controller.Token)
	r.GET("/me", controller.GetMeInfo)
	r.GET("/me/drive", controller.GetInfo)
	r.GET("/me/drive/root/*path", controller.GetChildren)
}
