package fiscobcos

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/user"
)

func TestFiscoBcosClient_RegisterUser(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	body := req.UserRegisterReqDataBody{
		UserID: "tutest04",
	}

	res, err := fiscobcosClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClient_EnrollUser(t *testing.T) {
	fabricClient := getFiscoBcosClient(t)

	body := req.UserEnrollReqDataBody{
		UserID: "tutest04",
	}

	res, err := fabricClient.EnrollUser(body)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(res)
}
