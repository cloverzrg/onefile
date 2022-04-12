package credential

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/cloverzrg/onefile/db"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/model"
	"golang.org/x/oauth2"
)

type AzureIdentityAuthenticationProvider struct {
	tokenSource oauth2.TokenSource
	UserId      string
}

func (a AzureIdentityAuthenticationProvider) GetToken(ctx context.Context, options policy.TokenRequestOptions) (*azcore.AccessToken, error) {
	token, err := a.tokenSource.Token()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &azcore.AccessToken{
		Token:     token.AccessToken,
		ExpiresOn: token.Expiry,
	}, nil
}

func NewAzureIdentityAuthenticationProvider(token *oauth2.Token) AzureIdentityAuthenticationProvider {
	return AzureIdentityAuthenticationProvider{
		tokenSource: oauthConfig.TokenSource(context.Background(), token),
	}
}

func NewAzureIdentityAuthenticationProviderByUserId(userId string) (a AzureIdentityAuthenticationProvider, err error) {
	token, err := model.GetTokenByUserId(db.DB, userId)
	if err != nil {
		logger.Error(err)
		return a, err
	}

	oauthToken := &oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
	return AzureIdentityAuthenticationProvider{
		tokenSource: oauthConfig.TokenSource(context.Background(), oauthToken),
		UserId:      userId,
	}, err
}
