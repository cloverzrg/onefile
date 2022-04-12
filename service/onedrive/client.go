package onedrive

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/cloverzrg/onefile/credential"
	"github.com/cloverzrg/onefile/db"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/model"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func getClientByUserId(userId string) (client *msgraphsdk.GraphServiceClient, err error) {
	authProvider, err := credential.NewAzureIdentityAuthenticationProviderByUserId(userId)
	if err != nil {
		logger.Errorf("Error authentication provider: %v\n", err)
		return nil, err
	}
	return getClientByAuthProvider(authProvider)
}

func getClientByToken(token *oauth2.Token) (client *msgraphsdk.GraphServiceClient, err error) {
	authProvider := credential.NewAzureIdentityAuthenticationProvider(token)
	return getClientByAuthProvider(authProvider)
}

func getClientByAuthProvider(authProvider azcore.TokenCredential) (client *msgraphsdk.GraphServiceClient, err error) {
	auth, err := a.NewAzureIdentityAuthenticationProvider(authProvider)
	if err != nil {
		logger.Errorf("Error authentication provider: %v\n", err)
		return nil, err
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		logger.Errorf("Error creating adapter: %v\n", err)
		return nil, err
	}

	client = msgraphsdk.NewGraphServiceClient(adapter)
	return client, err
}

func Callback(ctx context.Context, code string) (token *oauth2.Token, err error) {
	token, err = credential.GetTokenByCode(ctx, code)
	client, err := getClientByToken(token)
	if err != nil {
		logger.Errorf("Error creating client: %v\n", err)
		return nil, err
	}
	userInfo, err := client.Me().Get(nil)
	if err != nil {
		logger.Errorf("Error creating client: %v\n", err)
		return nil, err
	}
	bytes, _ := json.Marshal(userInfo)
	logger.Info(string(bytes))

	tokenByUserId, err := model.GetTokenByUserId(db.DB, *userInfo.GetId())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			token2 := &model.Token{
				UserId:       *userInfo.GetId(),
				AccessToken:  token.AccessToken,
				RefreshToken: token.RefreshToken,
				Expiry:       token.Expiry,
				TokenType:    token.TokenType,
			}
			token2, err = model.CreateToken(db.DB, token2)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
		} else {
			logger.Error(err)
			return token, err
		}
	}
	tokenByUserId.AccessToken = token.AccessToken
	tokenByUserId.RefreshToken = token.RefreshToken
	tokenByUserId.Expiry = token.Expiry
	tokenByUserId.TokenType = token.TokenType
	err = tokenByUserId.Save(db.DB)
	if err != nil {
		logger.Error(err)
		return token, err
	}
	return token, err
}