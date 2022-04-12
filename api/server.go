package api

import (
	"github.com/cloverzrg/onefile/api/middleware"
	"github.com/cloverzrg/onefile/logger"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithWriter(logger.Entry.Writer()))
	r.Use(middleware.GetSessionsMiddleware())
	r.Use(middleware.TokenMiddleware)
	SetRouter(r)
	return r.Run("[::]:80")
}
