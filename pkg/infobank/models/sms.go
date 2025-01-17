package models

import (
	"github.com/go-playground/validator/v10"
)

type SMS struct {
	From      string `json:"from" validate:"required"`
	To        string `json:"to" validate:"required"`
	Text      string `json:"text" validate:"required"`
	Ref       string `json:"ref,omitempty"`
	OriginCID string `json:"originCID,omitempty"`
}

func (m *SMS) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
