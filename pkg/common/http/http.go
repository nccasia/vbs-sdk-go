package http

import (
	"bytes"
	"net/http"

	"github.com/wonderivan/logger"
)

func SendPost(dataBytes []byte, url string) ([]byte, error) {
	var client *http.Client

	tr := new(http.Transport)
	tr.DisableKeepAlives = true
	client = &http.Client{
		//define the mechanism for a single HTTP request
		Transport: tr,
	}

	//invoke interface
	logger.Debug("request message：", string(dataBytes))
	response, err := client.Post(url, "application/json", bytes.NewReader(dataBytes))
	if err != nil {
		logger.Error("request failed：", err.Error())
		return nil, err
	}
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	//Get the response message data from the response object and read it
	allBytes := []byte{}
	bytes := make([]byte, response.ContentLength)
	i, _ := response.Body.Read(bytes)
	allBytes = append(allBytes, bytes[:i]...)

	for {
		i, _ = response.Body.Read(bytes)
		if i == 0 {
			break
		}
		allBytes = append(allBytes, bytes[:i]...)

	}
	logger.Debug("response message：", string(allBytes))
	return allBytes, nil
}
