package models

import "github.com/go-playground/validator/v10"

type Friendtalk struct {
	KakaoMsgType KakaoMsgType  `json:"msgType" validate:"required"`
	SenderKey    string        `json:"senderKey" validate:"required"`
	To           string        `json:"to" validate:"required"`
	Text         string        `json:"text" validate:"required"`
	Button       []KakaoButton `json:"button,omitempty"`
	ImgUrl       string        `json:"imgUrl,omitempty"`
	Fallback     *Fallback     `json:"fallback,omitempty"`
	Ref          string        `json:"ref,omitempty"`
}

func (m *Friendtalk) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
