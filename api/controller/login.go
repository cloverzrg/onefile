package controller

import (
	"fmt"
	"github.com/cloverzrg/onefile/config"
	"github.com/gin-gonic/gin"
	"strings"
)

func Index(c *gin.Context) {
	c.String(200, "here is onefile!")
}

func Login(c *gin.Context) {
	odConfig := config.Config.OneDrive
	scope := strings.ReplaceAll(odConfig.Scope, " ", "%20")
	url := fmt.Sprintf("https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s", odConfig.ClientId, odConfig.RedirectUri, scope)
	c.Redirect(302, url)
}

func Callback(c *gin.Context) {
	c.String(200, "here is onedrive callback!")
}
