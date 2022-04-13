package controller

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/cloverzrg/onefile/credential"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/service/onedrive"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "here is onefile!")
}

func Login(c *gin.Context) {
	c.Redirect(302, credential.GetAuthUri())
}

func Callback(c *gin.Context) {
	if c.Query("error") != "" {
		c.String(200, c.Query("error_description"))
		return
	}
	if c.Query("code") == "" {
		c.String(200, c.Query("code is empty"))
		return
	}
	code := c.Query("code")
	//session_state := c.Query("session_state")
	userId, err := onedrive.Callback(c, code)
	if err != nil {
		c.JSON(500, err)
		return
	}

	session := sessions.Default(c)
	session.Set("userId", userId)
	session.Save()
	c.JSON(200, "login success")
}

func Token(c *gin.Context) {
	authProvider, err := credential.NewAzureIdentityAuthenticationProviderByUserId(c.GetString("userId"))
	if err != nil {
		logger.Error(err)
		return
	}
	token, err := authProvider.GetToken(c, policy.TokenRequestOptions{})
	c.JSON(200, token)
}
