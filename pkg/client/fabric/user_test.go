package fabric

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/user"
)

func TestFabricClient_RegisterUser(t *testing.T) {
	fabricClient := getFabricClient(t)
	body := req.UserRegisterReqDataBody{
		UserID: "tutest01",
	}

	res, err := fabricClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClient_EnrollUser(t *testing.T) {

	fabricClient := getFabricClient(t)

	body := req.UserEnrollReqDataBody{
		UserID: "tutest01",
		Secret: "qROPhyDfxsdH",
	}

	res, err := fabricClient.EnrollUser(body)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(res)
}
