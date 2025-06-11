package models

import (
	"github.com/go-playground/validator/v10"
)

type BrandMessage struct {
	SenderKey              string                 `json:"senderKey" validate:"required"`
	SendType               KakaoSendType          `json:"sendType" validate:"required"`
	KakaoMsgType           KakaoMsgType           `json:"msgType" validate:"required"`
	To                     string                 `json:"to" validate:"omitempty"`
	Text                   string                 `json:"text,omitempty"`
	ImgUrl                 string                 `json:"imgUrl,omitempty"`
	Targeting              string                 `json:"targeting,omitempty"`
	TemplateCode           string                 `json:"templateCode,omitempty"`
	Button                 []KakaoButton          `json:"button,omitempty"`
	Fallback               *Fallback              `json:"fallback,omitempty"`
	GroupTagKey            string                 `json:"groupTagKey,omitempty"`
	Adult                  string                 `json:"adult,omitempty"`
	PushAlarm              string                 `json:"pushAlarm,omitempty"`
	MessageVariable        map[string]interface{} `json:"messageVariable,omitempty"`
	ButtonVariable         map[string]interface{} `json:"buttonVariable,omitempty"`
	CouponVariable         map[string]interface{} `json:"couponVariable,omitempty"`
	ImageVariable          map[string]interface{} `json:"imageVariable,omitempty"`
	VideoVariable          map[string]interface{} `json:"videoVariable,omitempty"`
	CommerceVariable       map[string]interface{} `json:"commerceVariable,omitempty"`
	CarouselVariable       map[string]interface{} `json:"carouselVariable,omitempty"`
	OriginCID              string                 `json:"originCID,omitempty"`
	UnsubscribePhoneNumber string                 `json:"unsubscribePhoneNumber,omitempty"`
	UnsubscribeAuthNumber  string                 `json:"unsubscribeAuthNumber,omitempty"`
	Ref                    string                 `json:"ref,omitempty"`
}

func (m *BrandMessage) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
