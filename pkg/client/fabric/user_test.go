package fabric

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/user"
)

func TestFabricClientRegisterUser(t *testing.T) {
	fabricClient := getFabricClient(t)
	body := req.UserRegisterReqDataBody{
		UserID: "tutest005",
	}

	res, err := fabricClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClientEnrollUser(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.UserEnrollReqDataBody{
		UserID: "tutest005",
		Secret: "FziGKWQvjJMo",
	}

	res, err := fabricClient.EnrollUser(body)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(res)
}
