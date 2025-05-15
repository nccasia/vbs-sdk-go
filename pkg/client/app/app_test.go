package app

import (
	"fmt"
	"testing"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/req"
)

func TestGetAppInfo(t *testing.T) {
	reqData := req.AppInfoReqData{}
	header := base.ReqHeader{
		UserCode: "UserCode1",
		AppCode:  "AppCode1",
	}

	api := "http://localhost:8889"
	reqData.Header = header
	res, err := GetAppInfo(&reqData, api, "")

	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		fmt.Println(res.Header.Msg)
	} else {
		fmt.Println(res.Body)
	}
}
