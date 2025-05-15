package app

import (
	"encoding/json"

	"github.com/nccasia/vbs-sdk-go/pkg/common/http"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/req"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/res"
)

// GetAppInfo query basic information of Dapp
func GetAppInfo(data *req.AppInfoReqData, baseApi string, cert string) (*res.AppInfoResData, error) {
	url := baseApi + "/api/app/info"

	reqBytes, _ := json.Marshal(data)
	resBytes, err := http.SendPost(reqBytes, url)

	if err != nil {
		return nil, err
	}

	resData := &res.AppInfoResData{}

	err = json.Unmarshal(resBytes, resData)

	if err != nil {
		return nil, err
	}

	return resData, nil
}
