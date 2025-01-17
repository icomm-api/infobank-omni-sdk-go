package models

import "github.com/go-playground/validator/v10"

type Omni struct {
	Destinations []Destination `json:"destinations" validate:"required,dive"`
	MessageFlow  []OmniMessage `json:"messageFlow,omitempty" validate:"required,dive"`
	MessageForm  string        `json:"messageForm,omitempty"`
	PaymentCode  string        `json:"paymentCode,omitempty"`
	Ref          string        `json:"ref,omitempty"`
}

type Destination struct {
	To           string            `json:"to" validate:"required"`
	ClientKey    string            `json:"clientKey,omitempty"`
	ReplaceWords map[string]string `json:"replaceWords,omitempty"`
}

type OmniMessage struct {
	SMS        *OmniSMS        `json:"sms,omitempty"`
	MMS        *OmniMMS        `json:"mms,omitempty"`
	RCS        *OmniRCS        `json:"rcs,omitempty"`
	AlimTalk   *OmniAlimTalk   `json:"alimtalk,omitempty"`
	FriendTalk *OmniFriendTalk `json:"friendtalk,omitempty"`
}

type OmniSMS struct {
	From      string `json:"from" validate:"required"`
	Text      string `json:"text" validate:"required"`
	Ttl       string `json:"ttl,omitempty"`
	OriginCID string `json:"originCID,omitempty"`
}

type OmniMMS struct {
	From      string   `json:"from" validate:"required"`
	Text      string   `json:"text,omitempty"`
	Title     string   `json:"title,omitempty"`
	FileKey   []string `json:"fileKey,omitempty"`
	Ttl       string   `json:"ttl,omitempty"`
	OriginCID string   `json:"originCID,omitempty"`
}

type OmniRCS struct {
	From         string     `json:"from" validate:"required"`
	FormatId     string     `json:"formatId" validate:"required"`
	Content      RcsContent `json:"content" validate:"required"`
	GroupId      string     `json:"groupId,omitempty"`
	ExpiryOption string     `json:"expiryOption,omitempty"`
	CopyAllowed  string     `json:"copyAllowed,omitempty"`
	Header       string     `json:"header,omitempty"`
	Footer       string     `json:"footer,omitempty"`
	Ttl          string     `json:"ttl,omitempty"`
	BrandId      string     `json:"brandId,omitempty"`
	BrandKey     string     `json:"brandKey,omitempty"`
	AgencyId     string     `json:"agencyId,omitempty"`
	AgencyKey    string     `json:"agencyKey,omitempty"`
	RcsData      string     `json:"rcsData,omitempty"`
}

type OmniAlimTalk struct {
	SenderKey    string              `json:"senderKey" validate:"required"`
	Text         string              `json:"text" validate:"required"`
	TemplateCode string              `json:"templateCode" validate:"required"`
	KakaoMsgType KakaoMsgType        `json:"msgType" validate:"required"`
	Title        string              `json:"title,omitempty"`
	Header       string              `json:"header,omitempty"`
	Attachment   *AlimTalkAttachment `json:"attachment,omitempty"`
	Price        string              `json:"price,omitempty"`
	CurrencyType string              `json:"currencyType,omitempty"`
	Supplement   *KakaoSuplement     `json:"supplement,omitempty"`
}

type OmniFriendTalk struct {
	SenderKey        string                `json:"senderKey" validate:"required"`
	KakaoMsgType     KakaoMsgType          `json:"msgType" validate:"required"`
	Text             string                `json:"text,omitempty"`
	Attachment       *FriendTalkAttachment `json:"attachment,omitempty"`
	AdFlag           string                `json:"adFlag,omitempty"`
	AddtionalContent string                `json:"addtionalContent,omitempty"`
	Adult            string                `json:"adult,omitempty"`
	Header           string                `json:"header,omitempty"`
	Carousel         *KakaoCarousel        `json:"carousel,omitempty"`
	GroupTagKey      string                `json:"groupTagKey,omitempty"`
	PushAlarm        string                `json:"pushAlarm,omitempty"`
}

func (m *Omni) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
