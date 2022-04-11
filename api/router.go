package api

import (
	"github.com/cloverzrg/onefile/api/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.GET("/login", controller.Login)

}
