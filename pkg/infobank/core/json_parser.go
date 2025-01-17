package core

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

func MarshalAndConvertStr(vo interface{}) (*string, error) {
	b, err := marshal(vo)
	if err != nil {
		return nil, err
	} else {
		s := string(b)
		return &s, nil
	}
}

func marshal(vo interface{}) ([]byte, error) {
	return json.Marshal(vo)
}

func Unmarshal(data []byte, vo interface{}) error {
	des := json.NewDecoder(bytes.NewReader(data))
	if err := des.Decode(vo); err != nil {
		return err
	} else {
		return nil
	}
}

func Valid(data []byte) bool {
	return json.Valid(data)
}

func ParseSimpleSendResponse(client *HttpClient, httpRes *http.Response) (*models.SimpleSendResponse, error) {
	apiRes, err := client.ParseHttpResponseBody(httpRes, new(models.SimpleSendResponse))
	if err != nil {
		return nil, err
	}

	res := apiRes.(*models.SimpleSendResponse)
	return res, nil
}

func ParseOmniSendResponse(client *HttpClient, httpRes *http.Response) (*models.OmniSendResponse, error) {
	apiRes, err := client.ParseHttpResponseBody(httpRes, new(models.OmniSendResponse))
	if err != nil {
		return nil, err
	}

	res := apiRes.(*models.OmniSendResponse)
	return res, nil
}

func ParseReportResponse(client *HttpClient, httpRes *http.Response) (*models.ReportResponse, error) {
	apiRes, err := client.ParseHttpResponseBody(httpRes, new(models.ReportResponse))
	if err != nil {
		return nil, err
	}

	res := apiRes.(*models.ReportResponse)
	return res, nil
}

func ParseFileResponse(client *HttpClient, httpRes *http.Response) (*models.FileResponse, error) {
	apiRes, err := client.ParseHttpResponseBody(httpRes, new(models.FileResponse))
	if err != nil {
		return nil, err
	}

	res := apiRes.(*models.FileResponse)
	return res, nil
}

func ParseFormResponse(client *HttpClient, httpRes *http.Response) (*models.FormResponse, error) {
	apiRes, err := client.ParseHttpResponseBody(httpRes, new(models.FormResponse))
	if err != nil {
		return nil, err
	}

	res := apiRes.(*models.FormResponse)
	return res, nil
}
