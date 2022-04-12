package credential

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"time"
)

type LocalTokenProvider struct {
	token string
}

func (a LocalTokenProvider) GetToken(ctx context.Context, options policy.TokenRequestOptions) (*azcore.AccessToken, error) {
	return &azcore.AccessToken{
		Token:     a.token,
		ExpiresOn: time.Now().Add(1 * time.Hour),
	}, nil
}

func NewLocalTokenProvider(token string) LocalTokenProvider {
	return LocalTokenProvider{
		token: token,
	}
}
