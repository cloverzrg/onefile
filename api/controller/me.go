package controller

import (
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/service/onedrive"
	"github.com/cloverzrg/onefile/util"
	"github.com/gin-gonic/gin"
)

func GetMeInfo(c *gin.Context) {
	client, err := onedrive.GetClientByUserId(c.GetString("userId"))
	if err != nil {
		logger.Error(err)
		return
	}
	userInfo, err := client.Me().Get(nil)
	if err != nil {
		logger.Error(err)
		return
	}
	res, err := util.ToJson(userInfo)
	if err != nil {
		logger.Error(err)
		c.AbortWithError(500, err)
		return
	}
	c.String(200, res)
}
