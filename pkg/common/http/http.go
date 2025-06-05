package http

import (
	"bytes"
	"fmt"
	"io"
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

	// Kiểm tra mã trạng thái HTTP
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		fmt.Printf("Error: %d %s, body: %s", response.StatusCode, response.Status, string(body))
		return nil, fmt.Errorf("unexpected status code: %d %s", response.StatusCode, response.Status)
	}

	// Đọc toàn bộ response body
	allBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	logger.Debug("response message：", string(allBytes))
	return allBytes, nil
}
