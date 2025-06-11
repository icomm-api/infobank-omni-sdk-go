package models

import "github.com/go-playground/validator/v10"

type Alimtalk struct {
	KakaoMsgType KakaoMsgType  `json:"msgType" validate:"required"`
	SenderKey    string        `json:"senderKey" validate:"required"`
	To           string        `json:"to" validate:"required"`
	Text         string        `json:"text" validate:"required"`
	TemplateCode string        `json:"templateCode" validate:"required"`
	Title        string        `json:"title,omitempty"`
	Header       string        `json:"header,omitempty"`
	Link         *KakaoLink    `json:"link,omitempty"`
	Button       []KakaoButton `json:"button,omitempty" validate:"dive"`
	Fallback     *Fallback     `json:"fallback,omitempty"`
	Ref          string        `json:"ref,omitempty"`
}

func (m *Alimtalk) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
