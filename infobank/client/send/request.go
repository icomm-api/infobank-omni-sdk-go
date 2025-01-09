package send

import (
	"bytes"
	"net/http"
	"omni_sdk_go/infobank"
	"omni_sdk_go/infobank/client/core"
)

const (
	contentType = "application/json"
)

func sendMessage(client *core.HttpClient, subPath string, authorization string, req interface{}) (*http.Response, error) {

	url := infobank.RemoteAddress + subPath

	header := make(map[string]string)
	header["Content-Type"] = contentType
	header["Authorization"] = authorization

	parsedBody, parseErr := core.MarshalAndConvertStr(req)
	if parseErr != nil {
		return nil, *parseErr
	}

	body := &bytes.Buffer{}
	if parsedBody != nil {
		body.WriteString(*parsedBody)
	}

	httpRes, httpErr := client.Request(http.MethodPost, url, header, body)
	if httpErr != nil {
		return nil, *httpErr
	}

	return httpRes, nil
}
