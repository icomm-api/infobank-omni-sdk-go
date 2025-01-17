package rcs

import (
	"github.com/go-playground/validator/v10"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	rcsSubPath = "/v1/send/rcs"
)

type RcsSender struct {
	Client *core.HttpClient
}

func (sender *RcsSender) SendMessage(req *models.RCS) (*models.SimpleSendResponse, error) {
	err := req.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, err
		}
	}

	response, err := core.SendMessage(sender.Client, rcsSubPath, req)
	if err != nil {
		return nil, err
	}
	return core.ParseSimpleSendResponse(sender.Client, response)
}
