# infobank-omni-sdk-go v1.0.0
---
[![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)]() [![Migration Guide](https://img.shields.io/badge/Migration-Guide-blue)]() [![API Reference](https://img.shields.io/badge/api-reference-blue.svg)]() [![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)]()


infobank omni api sdk go 입니다.

바로가기 :
- [OMNI API](https://infobank-guide.gitbook.io/omni_api)
- [시작하기](#시작하기-getting-started)
- [사용법](#사용법-usage)
  - [토큰 발급](#토큰-발급)
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

### 토큰 발급
```go
import (
    "fmt"
    "github.com/icomm-api/infobank-omni-sdk-go/infobank/client/auth"
)

func main() {

    clientId := "test"
    password := "1234"
    
    authClient := auth.NewAuthenticator(nil, clientId, password)
    err := authClient.Init()
    if err != nil {
        fmt.Println("err :", err)
    } else {
        token := authClient.GetToken()
        fmt.Println(*token)
    }
}
```


### File 업로드
```go
import (
    "fmt"
    "github.com/icomm-api/infobank-omni-sdk-go/infobank/client/regist"
)

func main() {
    fileUploader := regist.NewFileUploader("input-your-token", nil)
    serviceType := FILE_SERVICE_TYPE_MMS
    response, err := fileUploader.UploadFile(&serviceType, nil, "D:\\", "2.jpg")

    if err != nil {
        fmt.Println("err :", err)
	} else {
        fmt.Println(response)
    }
}
```


### FORM 등록
```go
import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/icomm-api/infobank-omni-sdk-go/infobank/client/regist"
    "github.com/icomm-api/infobank-omni-sdk-go/infobank/client/send"
)

func mian() {
    formRegister := regist.NewFormRegister("input-your-token", nil)
    msg := send.NewOmniSmsBuilder().Text("text").From("01012345678").From("01012345678").Build()
    form := NewMessageFormBuilder().Message(msg).Build()
    apiResponse, err := formRegister.RegisterForm(form)
    
	if err != nil {
        fmt.Println("err :", err)
    } else {
        if apiResponse.Code != "A000" {
            fmt.Println("err :", err)
        } else {
            fmt.Println(apiResponse.Result, apiResponse.Code, *apiResponse.Data.FormId, *apiResponse.Data.Expired)
        }
    }
}
```

### 전송
```go
package send

import (
	"errors"
	"fmt"
	"github.com/icomm-api/infobank-omni-sdk-go/infobank/client/send"
)

func main() {

	smsSender := send.NewSmsSender("input-your-token", nil)
	message := send.NewSmsBuilder().From("0310000000").To("01012345618").Text("omni-sdk-go text").Ref("omni-sdk-go ref").OriginCID("0000").Build()

	apiResponse, err := smsSender.SendMessage(&message)

	if err != nil {
		fmt.Println("err :", err)
		return
	}

	if apiResponse.Code != "A000" {
		fmt.Println("err : ", errors.New(apiResponse.Result))
		return
	}
	fmt.Println(*apiResponse)
}
```

### 리포트
```go
import (
    "encoding/json"
    "fmt"
    "github.com/icomm-api/infobank-omni-sdk-go/infobank/client/report"
)

func main() {
    reportManager := report.NewReportManager("input-your-token", nil)
    apiResponse, err := reportManager.InquiryReport("20230619081619396PRDR1SM92760200")
    if err != nil {
        fmt.Println("err :", err)
    } else {
        data, _ := json.Marshal(apiResponse.Data)
        fmt.Println(apiResponse.Code, apiResponse.Data, string(data))
    }
}
```

## 문의 (Contact)
본 문서와 관련된 기술 문의는 아래 메일 주소로 연락 바랍니다. 😄

[support@infobank.net](support@infobank.net)








