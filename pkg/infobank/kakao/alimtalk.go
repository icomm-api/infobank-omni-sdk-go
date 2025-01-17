package kakao

import (
	"github.com/go-playground/validator/v10"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	alimtalkSubPath = "/v1/send/alimtalk"
)

type AlimtalkSender struct {
	Client *core.HttpClient
}

func (sender *AlimtalkSender) SendMessage(req *models.Alimtalk) (*models.SimpleSendResponse, error) {
	err := req.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, err
		}
	}

	response, err := core.SendMessage(sender.Client, alimtalkSubPath, req)
	if err != nil {
		return nil, err
	}
	return core.ParseSimpleSendResponse(sender.Client, response)
}
