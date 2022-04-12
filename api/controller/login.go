package controller

import (
	"github.com/cloverzrg/onefile/credential"
	"github.com/cloverzrg/onefile/service/onedrive"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "here is onefile!")
}

func Login(c *gin.Context) {
	c.Redirect(302, credential.GetAuthUri())
}

func Callback(c *gin.Context) {
	code := c.Query("code")
	//session_state := c.Query("session_state")
	token, err := onedrive.Callback(c, code)
	if err != nil {
		c.JSON(500, "500")
		return
	}
	c.JSON(200, token)
}
