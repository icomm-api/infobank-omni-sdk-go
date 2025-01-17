package models

import (
	"github.com/go-playground/validator/v10"
)

type RCS struct {
	From         string      `json:"from" validate:"required"`
	To           string      `json:"to" validate:"required"`
	FormatId     string      `json:"formatId" validate:"required"`
	BrandKey     string      `json:"brandKey,omitempty"`
	BrandId      string      `json:"brandId,omitempty"`
	Content      *RcsContent `json:"content" validate:"required"`
	ExpiryOption string      `json:"expiryOption,omitempty"`
	Header       string      `json:"header,omitempty"`
	Footer       string      `json:"footer,omitempty"`
	Fallback     *Fallback   `json:"fallback,omitempty"`
	Ref          string      `json:"ref,omitempty"`
	RcsData      string      `json:"rcsData,omitempty"`
}

type RcsContent struct {
	StandAlone *RcsStandAlone         `json:"standalone,omitempty"`
	Carousel   []RcsCarousel          `json:"carousel,omitempty"`
	Template   map[string]interface{} `json:"template,omitempty"`
}

type RcsStandAlone struct {
	Title      string          `json:"title,omitempty"`
	Text       string          `json:"text,omitempty"`
	Media      string          `json:"media,omitempty"`
	MediaUrl   string          `json:"mediaUrl,omitempty"`
	Button     []RcsButton     `json:"button,omitempty" validate:"dive"`
	SubContent []RcsSubContent `json:"subContent,omitempty"`
}

type RcsCarousel struct {
	Title    string      `json:"title,omitempty"`
	Text     string      `json:"text,omitempty"`
	Media    string      `json:"media,omitempty"`
	MediaUrl string      `json:"mediaUrl,omitempty"`
	FileKey  string      `json:"fileKey,omitempty"`
	Button   []RcsButton `json:"button,omitempty" validate:"dive"`
}

type RcsButton struct {
	Type        RcsButtonType `json:"type," validate:"required"`
	Name        string        `json:"name,omitempty" validate:"required"`
	Url         string        `json:"url,omitempty"`
	Label       string        `json:"label,omitempty"`
	Latitude    string        `json:"latitude,omitempty"`
	Longitude   string        `json:"longitude,omitempty"`
	FallbackUrl string        `json:"fallbackUrl,omitempty"`
	Query       string        `json:"query,omitempty"`
	StartTime   string        `json:"startTime,omitempty"`
	EndTime     string        `json:"endTime,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Text        string        `json:"text,omitempty"`
	PhoneNumber string        `json:"phoneNumber,omitempty"`
}

type RcsSubContent struct {
	SubTitle    string `json:"subTitle,omitempty"`
	SubDesc     string `json:"subDesc,omitempty"`
	SubMedia    string `json:"subMedia,omitempty"`
	SubMediaUrl string `json:"subMediaUrl,omitempty"`
}

func (m *RCS) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
