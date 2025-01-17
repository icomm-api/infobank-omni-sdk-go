package models

import "github.com/go-playground/validator/v10"

type MMS struct {
	From      string   `json:"from" validate:"required"`
	To        string   `json:"to" validate:"required"`
	Text      string   `json:"text,omitempty"`
	Title     string   `json:"title,omitempty"`
	FileKey   []string `json:"fileKey,omitempty"`
	Ref       string   `json:"ref,omitempty"`
	OriginCID string   `json:"originCID,omitempty"`
}

func (m *MMS) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
