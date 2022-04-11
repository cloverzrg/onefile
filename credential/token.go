package credential

import (
	"context"
	"github.com/cloverzrg/onefile/config"
	"github.com/cloverzrg/onefile/db"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
	"gorm.io/gorm"
)

var oauthConfig oauth2.Config

func init() {
	oauthConfig = oauth2.Config{
		ClientID:     config.Config.OneDrive.ClientId,
		ClientSecret: config.Config.OneDrive.ClientSecret,
		Endpoint:     microsoft.AzureADEndpoint(""),
		RedirectURL:  config.Config.OneDrive.RedirectUri,
		Scopes:       config.Config.OneDrive.Scope,
	}
}

func GetAuthUri() string {
	return oauthConfig.AuthCodeURL("")
}

func GetToken(ctx context.Context, code string) (token *oauth2.Token, err error) {
	token, err = oauthConfig.Exchange(ctx, code)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	token2 := model.Token{
		Model:        gorm.Model{},
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
	err = db.DB.Save(&token2).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return token, err
}
