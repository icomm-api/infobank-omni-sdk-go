package sms

import (
	"github.com/go-playground/validator/v10"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	smsSubPath = "/v1/send/sms"
)

type SmsSender struct {
	Auth   core.Authenticator
	Client *core.HttpClient
}

func (sender *SmsSender) SendMessage(req models.SMS) (*models.SimpleSendResponse, error) {
	err := req.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, err
		}
	}

	if response, err := core.SendMessage(sender.Client, smsSubPath, req); err != nil {
		return nil, err
	} else {
		return core.ParseSimpleSendResponse(sender.Client, response)
	}
}
