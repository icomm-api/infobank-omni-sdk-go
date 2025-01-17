package international

import (
	"github.com/go-playground/validator/v10"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	internationalSmsSubPath = "/v1/send/international"
	regex                   = `^\+[0-9]+$`
)

type SmsSender struct {
	Client *core.HttpClient
}

func (sender *SmsSender) SendMessage(req *models.SMS) (*models.SimpleSendResponse, error) {
	err := req.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, err
		}
	}

	response, err := core.SendMessage(sender.Client, internationalSmsSubPath, req)
	if err != nil {
		return nil, err
	}
	return core.ParseSimpleSendResponse(sender.Client, response)
}
