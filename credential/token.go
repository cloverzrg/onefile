package credential

import (
	"context"
	"github.com/cloverzrg/onefile/config"
	"github.com/cloverzrg/onefile/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var oauthConfig oauth2.Config

func init() {
	oauthConfig = oauth2.Config{
		ClientID:     config.Config.OneDrive.ClientId,
		ClientSecret: config.Config.OneDrive.ClientSecret,
		Endpoint:     microsoft.AzureADEndpoint("common"),
		RedirectURL:  config.Config.Baseurl + config.Config.OneDrive.RedirectUri,
		Scopes:       config.Config.OneDrive.Scope,
	}
}

func GetAuthUri() string {
	return oauthConfig.AuthCodeURL("")
}

func GetTokenByCode(ctx context.Context, code string) (token *oauth2.Token, err error) {
	token, err = oauthConfig.Exchange(ctx, code)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return token, err
}
