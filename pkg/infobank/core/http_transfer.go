package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	apiHttpStatusCode = map[int]bool{
		http.StatusOK:                   true,
		http.StatusBadRequest:           true,
		http.StatusUnauthorized:         true,
		http.StatusForbidden:            true,
		http.StatusNotFound:             true,
		http.StatusUnsupportedMediaType: true,
		http.StatusTooManyRequests:      true,
		http.StatusInternalServerError:  true,
		http.StatusServiceUnavailable:   true,
	}
)

type HttpClient struct {
	HttpClient *http.Client
}

func (c *HttpClient) Request(method string, url string, header map[string]string, data *bytes.Buffer) (*http.Response, error) {
	var request *http.Request
	var err error

	if data == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, data)
	}

	if err != nil {
		return nil, err
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if _, exists := apiHttpStatusCode[response.StatusCode]; !exists {
		errorMessage := fmt.Sprintf("httpStatusCode := %d, msg := %s", response.StatusCode, response.Status)
		return nil, errors.New(errorMessage)
	}

	return response, nil
}

func (c *HttpClient) ParseHttpResponseBody(res *http.Response, vo interface{}) (interface{}, error) {
	defer res.Body.Close()
	body, responseReadErr := io.ReadAll(res.Body)
	if responseReadErr != nil {
		return nil, responseReadErr
	}

	if len(body) == 0 {
		return nil, errors.New("response body is empty")
	}

	if marshalErr := Unmarshal(body, &vo); marshalErr != nil {
		return nil, marshalErr
	}

	return vo, nil
}

const (
	contentType = "application/json"
)

func SendMessage(client *HttpClient, subPath string, req interface{}) (*http.Response, error) {

	url := OmniUrl + subPath
	token, err := VerifyInValidToken()
	if err != nil {
		return nil, err
	}

	header := map[string]string{
		"Content-Type":  contentType,
		"Authorization": *token,
	}

	parsedBody, err := MarshalAndConvertStr(req)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	if parsedBody != nil {
		body.WriteString(*parsedBody)
	}

	response, err := client.Request(http.MethodPost, url, header, body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ManageReport(client *HttpClient, method string, subPath string, reportId *string) (*http.Response, error) {
	token, err := VerifyInValidToken()
	if err != nil {
		return nil, err
	}

	url := OmniUrl + subPath
	header := map[string]string{
		"Authorization": *token,
		"Content-Type":  contentType,
	}

	if reportId != nil {
		url += "/" + *reportId
	}

	response, err := client.Request(method, url, header, nil)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func UploadFile(client *HttpClient, subPath string, contentType string, body *bytes.Buffer) (*http.Response, error) {
	url := OmniUrl + subPath
	token, err := VerifyInValidToken()
	if err != nil {
		return nil, err
	}

	header := map[string]string{
		"Content-Type":  contentType,
		"Authorization": *token,
	}

	response, err := client.Request(http.MethodPost, url, header, body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ManageForm(client *HttpClient, method string, subPath string, req interface{}) (*http.Response, error) {
	url := OmniUrl + subPath
	token, err := VerifyInValidToken()
	if err != nil {
		return nil, err
	}
	var body *bytes.Buffer
	header := map[string]string{
		"Authorization": *token,
	}

	if method == http.MethodPost || method == http.MethodPut {
		header["Content-Type"] = contentType

		parsedBody, err := MarshalAndConvertStr(req)
		if err != nil {
			return nil, err
		}
		body = new(bytes.Buffer)
		if parsedBody != nil {
			body.WriteString(*parsedBody)
		}
	}

	response, err := client.Request(method, url, header, body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func InquiryForm(client *HttpClient, method string, subPath string, formId *string) (*http.Response, error) {
	url := OmniUrl + subPath
	token, err := VerifyInValidToken()
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"Authorization": *token,
	}

	url = url + "/" + *formId

	header["Content-Type"] = contentType

	response, err := client.Request(method, url, header, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func DeleteForm(client *HttpClient, method string, subPath string, formId *string) (*http.Response, error) {
	url := OmniUrl + subPath
	token, err := VerifyInValidToken()
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"Authorization": *token,
	}

	url = url + "/" + *formId

	header["Content-Type"] = contentType

	response, err := client.Request(method, url, header, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}
