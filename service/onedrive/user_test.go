package onedrive

import (
	"fmt"
	"github.com/cloverzrg/onefile/credential"
	"github.com/cloverzrg/onefile/util"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children"
	"os"
	"testing"
)

var token = os.Getenv("token")

func TestGetUserInfo(t *testing.T) {

	client, _ := getClientByAuthProvider(credential.NewLocalTokenProvider(token))
	userInfo, err := client.Me().Get(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(userInfo.GetId())
	json2, err := util.ToJson(userInfo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(json2)

}

func TestGetDrive(t *testing.T) {
	// https://docs.microsoft.com/en-us/graph/api/drive-get?view=graph-rest-1.0&tabs=go
	client, _ := getClientByAuthProvider(credential.NewLocalTokenProvider(token))
	// userId := "user-id"
	// result, err := graphClient.UsersById(&userId).Drive().Get(nil)

	// siteId := "site-id"
	// result, err := graphClient.SitesById(&siteId).Drive().Get(nil)

	// 基础信息
	//driveInfo, err := client.Me().Drive().Get(nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(*driveInfo.GetId())

	// 子目录
	itemable, err := client.Me().DrivesById("b!vGVCI3uy2EiyqN-yjdplyjZcLSwyU-FIkxN39hF44mZUOoZH_WjvR5lM-NvSAIb-").Root().Children().Get(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(itemable.GetValue())

	// https://docs.microsoft.com/en-us/graph/api/driveitem-list-children?view=graph-rest-1.0&tabs=http
	path := "/程序自动创建的文件夹"
	targetUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/%s:/children", path)

	auth, err := a.NewAzureIdentityAuthenticationProvider(credential.NewLocalTokenProvider(token))
	if err != nil {
		fmt.Printf("Error authenticating: %v\n", err)
		return
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		fmt.Printf("Error creating adapter: %v\n", err)
		return
	}

	client2 := children.NewChildrenRequestBuilder(targetUrl, adapter)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	result, err := client2.Get(nil)
	if err != nil {
		fmt.Printf("Error retrieving resource: %v\n", err)
	}
	fmt.Printf("result = %+v\n", result)
	json2, err := util.ToJson(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(json2)
}
