package regist

import (
	"net/http"

	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	formSubPath = "/v1/form"
	contentType = "application/json"
)

type FormRegister struct {
	Client *core.HttpClient
}

func (f *FormRegister) RegisterForm(req models.MessageForm) (*models.FormResponse, error) {
	response, err := core.ManageForm(f.Client, http.MethodPost, formSubPath, req)
	if err != nil {
		return nil, err
	}
	return core.ParseFormResponse(f.Client, response)
}

func (f *FormRegister) InquiryForm(formId string) (*models.FormResponse, error) {
	response, err := core.InquiryForm(f.Client, http.MethodGet, formSubPath, &formId)
	if err != nil {
		return nil, err
	}
	return core.ParseFormResponse(f.Client, response)
}

func (f *FormRegister) DeleteForm(formId string) (*models.FormResponse, error) {
	response, err := core.DeleteForm(f.Client, http.MethodDelete, formSubPath, &formId)
	if err != nil {
		return nil, err
	}
	return core.ParseFormResponse(f.Client, response)
}
