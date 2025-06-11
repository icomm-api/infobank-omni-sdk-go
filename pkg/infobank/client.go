package infobank

import (
	"errors"
	"net/http"

	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/international"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/kakao"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/mms"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/omni"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/rcs"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/regist"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/report"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/sms"
)

type OmniClient struct {
	SMS              sms.SmsSender
	MMS              mms.MmsSender
	InternationalSMS international.SmsSender
	RCS              rcs.RcsSender
	AlimTalk         kakao.AlimtalkSender
	FriendTalk       kakao.FriendtalkSender
	BrandMessage     kakao.BrandMessageSender
	Omni             omni.OmniSender
	Report           report.ReportManager
	File             regist.FileUploader
	Form             regist.FormRegister
}

func NewOmniClient(url string, clientId string, password string) (*OmniClient, error) {
	if url == "" || clientId == "" || password == "" {
		return nil, errors.New("url, clientId, and password must not be empty")
	}
	c := &OmniClient{}

	core.Init(url, clientId, password)

	c.SMS = sms.SmsSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.MMS = mms.MmsSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.InternationalSMS = international.SmsSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.RCS = rcs.RcsSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.AlimTalk = kakao.AlimtalkSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.FriendTalk = kakao.FriendtalkSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.BrandMessage = kakao.BrandMessageSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.Omni = omni.OmniSender{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.Report = report.ReportManager{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.File = regist.FileUploader{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	c.Form = regist.FormRegister{
		Client: &core.HttpClient{
			HttpClient: &http.Client{},
		},
	}

	return c, nil
}
