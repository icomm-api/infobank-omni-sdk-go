# infobank-omni-sdk-go v1.1.0
---
[![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)]() [![Migration Guide](https://img.shields.io/badge/Migration-Guide-blue)]() [![API Reference](https://img.shields.io/badge/api-reference-blue.svg)]() [![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)]()


infobank omni api sdk go 입니다.

바로가기 :
- [OMNI API](https://infobank-guide.gitbook.io/omni_api)
- [시작하기](#시작하기-getting-started)
- [사용법](#사용법-usage)
  - [파일 업로드](#파일-업로드)
  - [FORM 등록](#form-등록)
  - [전송](#전송)
  - [리포트](#리포트)
- [문의](#문의)

------------
## 시작하기 (Getting Started)

```bash
go get github.com/icomm-api/infobank-omni-sdk-go
```


## 사용법 (Usage)

### File 업로드
```go
import (
    "github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank"
    "github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

func main() {
    omniClient, err := infobank.NewOmniClient(Url, clientId, clientPassword)
    
    serviceType := models.FILE_SERVICE_TYPE_MMS
    msgType := ""
	
    res, err := omniClient.File.UploadFile(serviceType, msgType, "<이미지 경로>")
}
```


### FORM 등록
```go
import (
    "github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank"
    "github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

func mian() {
    omniClient, err := infobank.NewOmniClient(Url, clientId, clientPassword)
    
    message := models.MessageForm{
		MessageForm: []models.OmniMessage{
			{
				SMS: &models.OmniSMS{
					From: "<발신번호>",
					Text: "<Text>",
				},
			},
		},
	}

    res, err := omniClient.Form.RegisterForm(message)
}
```

### 전송
```go
package send

import (
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank"
    "github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

func main() {
	omniClient, err := infobank.NewOmniClient(Url, clientId, clientPassword)
	
    mms := models.MMS{
		From:    "<발신번호>",
		To:      "<수신번호",
		Text:    "<Text>",
	}

	res, err := omniClient.SMS.SendMessage(&mms)
}
```

### 리포트
```go
import (
    "github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank"
)

func main() {
    omniClient, err := infobank.NewOmniClient(Url, clientId, clientPassword)
    msgKey := "<msgKey>"
	res, err := omniClient.Report.InquiryReport(msgKey)
}
```

## 문의 (Contact)
본 문서와 관련된 기술 문의는 아래 메일 주소로 연락 바랍니다. 😄

[support@infobank.net](support@infobank.net)








