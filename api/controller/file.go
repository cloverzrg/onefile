package controller

import (
	"fmt"
	"github.com/cloverzrg/onefile/credential"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/util"
	"github.com/gin-gonic/gin"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children"
	"strings"
)

func GetFile(c *gin.Context) {
	path := c.Param("path")
	path = strings.Trim(path, "/")
	path = "/" + path
	logger.Info("path:", path)
	targetUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/%s:/children", path)
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
