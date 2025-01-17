package omni

import (
	"github.com/go-playground/validator/v10"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	omniSubPath = "/v1/send/omni"
)

type OmniSender struct {
	Client *core.HttpClient
}

func (sender *OmniSender) SendMessage(req *models.Omni) (*models.OmniSendResponse, error) {
	err := req.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, err
		}
	}

	response, err := core.SendMessage(sender.Client, omniSubPath, req)
	if err != nil {
		return nil, err
	}
	return core.ParseOmniSendResponse(sender.Client, response)
}
