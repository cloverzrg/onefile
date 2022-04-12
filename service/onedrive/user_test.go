package onedrive

import (
	"github.com/cloverzrg/onefile/credential"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	client, _ := getClientByAuthProvider(credential.NewLocalTokenProvider("..--"))
	userInfo, err := client.Me().Get(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(userInfo.GetId())
}
