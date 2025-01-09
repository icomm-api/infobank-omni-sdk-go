package send

import (
	"bytes"
	"net/http"

	"github.com/icomm-api/infobank-omni-sdk-go.git/infobank/client/core"

	"github.com/icomm-api/infobank-omni-sdk-go.git/infobank"
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
