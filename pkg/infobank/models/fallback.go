package models

type Fallback struct {
	Type      FallbackServiceType `json:"type" validate:"required"`
	Title     string              `json:"title,omitempty"`
	Text      string              `json:"text,omitempty"`
	FileKey   []string            `json:"fileKey,omitempty"`
	From      string              `json:"from,omitempty"`
	OriginCID string              `json:"originCID,omitempty"`
}
