package kakao

import (
	"github.com/go-playground/validator/v10"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	friendTalkSubPath = "/v1/send/friendtalk"
)

type FriendtalkSender struct {
	Client *core.HttpClient
}

func (sender *FriendtalkSender) SendMessage(req *models.Friendtalk) (*models.SimpleSendResponse, error) {
	err := req.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, err
		}
	}

	response, err := core.SendMessage(sender.Client, friendTalkSubPath, req)
	if err != nil {
		return nil, err
	}
	return core.ParseSimpleSendResponse(sender.Client, response)
}
