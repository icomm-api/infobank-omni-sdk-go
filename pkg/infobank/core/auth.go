package core

import (
	"errors"
	"net/http"
	"time"
)

const (
	subPath    = "/v1/auth/token"
	timeFormat = "2006-01-02T15:04:05+07:00"
)

var auth *Authenticator
var OmniUrl string

type TokenResponse struct {
	Code   string     `json:"code"`
	Result string     `json:"result"`
	Data   *TokenData `json:"data,omitempty"`
}

type TokenData struct {
	Token   string `json:"token"`
	Schema  string `json:"schema"`
	Expired string `json:"expired"`
}

type Authenticator struct {
	ClientId string
	Password string
	Client   *HttpClient

	authorization *TokenData
}

func Init(address string, clientId string, password string) {

	OmniUrl = address
	auth = &Authenticator{
		ClientId: clientId,
		Password: password,
		Client: &HttpClient{
			HttpClient: &http.Client{},
		},
	}
}

func VerifyInValidToken() (*string, error) {
	isTokenExpired := auth.verifyInValidToken()
	if isTokenExpired {
		response, err := auth.requestToken()
		if err != nil {
			return nil, err
		}

		if response.Code != "A000" {
			err = errors.New(response.Result)
			return nil, err
		}

		auth.authorization = response.Data

	}
	token := auth.GetToken()
	return token, nil
}

func (auth *Authenticator) GetToken() *string {
	if auth.authorization == nil {
		return nil
	} else {
		token := auth.authorization.Schema + " " + auth.authorization.Token
		return &token
	}
}

func (auth *Authenticator) verifyInValidToken() bool {

	currentTime := time.Now().Add(time.Minute * 10)
	if auth.authorization == nil {
		return true
	}

	expiredAt, err := time.Parse(timeFormat, auth.authorization.Expired)
	if err != nil {
		return true
	}

	if currentTime.After(expiredAt) {
		return true
	}

	return false
}

func (auth *Authenticator) requestToken() (*TokenResponse, error) {
	url := OmniUrl + subPath

	header := map[string]string{
		"X-IB-Client-Id":     auth.ClientId,
		"X-IB-Client-Passwd": auth.Password,
		"Content-Type":       "application/json",
	}

	response, err := auth.Client.Request(http.MethodPost, url, header, nil)
	if err != nil {
		return nil, err
	}

	apiRes, parseErr := auth.Client.ParseHttpResponseBody(response, new(TokenResponse))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*TokenResponse)
	return res, nil
}
