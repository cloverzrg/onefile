package controller

import (
	"fmt"
	"github.com/cloverzrg/onefile/credential"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/service/onedrive"
	"github.com/cloverzrg/onefile/util"
	"github.com/gin-gonic/gin"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children"
	"strings"
)

func GetChildren(c *gin.Context) {
	path := c.Param("path")
	path = strings.Trim(path, "/")
	path = "/" + path
	logger.Info("path:", path)
	targetUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/%s:/children?$select=id,name", path)
	authenticationProvider, err := credential.NewAzureIdentityAuthenticationProviderByUserId(c.GetString("userId"))
	if err != nil {
		logger.Infof("Error creating adapter: %v\n", err)
		return
	}
	auth, err := a.NewAzureIdentityAuthenticationProvider(authenticationProvider)
	if err != nil {
		logger.Infof("Error authenticating: %v\n", err)
		return
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactory(auth, util.NewJsonParseNodeFactory(), util.NewJsonSerializationWriterFactory())
	if err != nil {
		logger.Infof("Error creating adapter: %v\n", err)
		return
	}

	client2 := children.NewChildrenRequestBuilder(targetUrl, adapter)
	if err != nil {
		logger.Infof("Error creating client: %v\n", err)
		return
	}

	result, err := client2.Get(nil)
	if err != nil {
		logger.Infof("Error retrieving resource: %v\n", err)
		return
	}
	json2, err := util.ToJson(result)
	if err != nil {
		logger.Info(err)
		return
	}
	c.String(200, json2)
}

func GetInfo(c *gin.Context) {
	client, err := onedrive.GetClientByUserId(c.GetString("userId"))
	if err != nil {
		logger.Error(err)
		return
	}
	driveInfo, err := client.Me().Drive().Get(nil)
	if err != nil {
		logger.Error(err)
		c.JSON(500, err.Error())
		return
	}
	json2, err := util.ToJson(driveInfo)
	if err != nil {
		logger.Error(err)
		c.JSON(500, err.Error())
		return
	}
	c.String(200, json2)
}
