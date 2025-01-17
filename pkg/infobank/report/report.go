package report

import (
	"net/http"

	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	pollingSubPath = "/v1/report/polling"
	inquirySubPath = "/v1/report/inquiry"
)

type ReportManager struct {
	Client *core.HttpClient
}

func (r *ReportManager) PollingReport() (*models.ReportResponse, error) {
	response, err := core.ManageReport(r.Client, http.MethodGet, pollingSubPath, nil)
	if err != nil {
		return nil, err
	}

	return core.ParseReportResponse(r.Client, response)
}

func (r *ReportManager) DeleteReport(reportId string) (*models.ReportResponse, error) {
	response, err := core.ManageReport(r.Client, http.MethodDelete, pollingSubPath, &reportId)
	if err != nil {
		return nil, err
	}
	return core.ParseReportResponse(r.Client, response)
}

func (r *ReportManager) InquiryReport(reportId string) (*models.ReportResponse, error) {
	response, err := core.ManageReport(r.Client, http.MethodGet, inquirySubPath, &reportId)
	if err != nil {
		return nil, err
	}
	return core.ParseReportResponse(r.Client, response)
}
