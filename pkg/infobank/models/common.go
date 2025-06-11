package models

type FileServiceTypeEnum string

const (
	FILE_SERVICE_TYPE_MMS        FileServiceTypeEnum = "MMS"
	FILE_SERVICE_TYPE_RCS        FileServiceTypeEnum = "RCS"
	FILE_SERVICE_TYPE_FRIENDTALK FileServiceTypeEnum = "FRIENDTALK"
)

type FileMsgTypeEnum string

const (
	FILE_MSG_TYPE_FRIENDTALK_IMAGE                   FileMsgTypeEnum = "FI"
	FILE_MSG_TYPE_FRIENDTALK_WIDE_IMAGE              FileMsgTypeEnum = "FW"
	FILE_MSG_TYPE_FRIENDTALK_ITEMLIST_IMAGE          FileMsgTypeEnum = "FL"
	FILE_MSG_TYPE_FRIENDTALK_CAROUSEL_IMAGE          FileMsgTypeEnum = "FC"
	FILE_MSG_TYPE_FRIENDTALK_CAROUSEL_COMMERCE_IMAGE FileMsgTypeEnum = "FA"
	FILE_MSG_TYPE_BRANDMESSAGE_DEFAULT               FileMsgTypeEnum = "default"
	FILE_MSG_TYPE_BRANDMESSAGE_WIDE                  FileMsgTypeEnum = "wide"
	FILE_MSG_TYPE_BRANDMESSAGE_WIDEITEMLIST          FileMsgTypeEnum = "wideItemList"
	FILE_MSG_TYPE_BRANDMESSAGE_CAROUSELFEED          FileMsgTypeEnum = "carouselFeed"
	FILE_MSG_TYPE_BRANDMESSAGE_CAROUSELCOMMERCE      FileMsgTypeEnum = "carouselCommerce "
)

type FallbackServiceType string

const (
	FALLBACK_SMS FallbackServiceType = "SMS"
	FALLBACK_MMS FallbackServiceType = "MMS"
)

type KakaoMsgType string

const (
	MSGTYPE_ALIMTALK                       KakaoMsgType = "AT"
	MSGTYPE_ALIMTALK_IMAGE                 KakaoMsgType = "AI"
	MSGTYPE_FRIENDTALK                     KakaoMsgType = "FT"
	MSGTYPE_FRIENDTALK_IMAGE               KakaoMsgType = "FI"
	MSGTYPE_FRIENDTALK_WIDE_IMAGE          KakaoMsgType = "FW"
	MSGTYPE_FRIENDTALK_CAROUSEL            KakaoMsgType = "FC"
	MSGTYPE_FRIENDTALK_ITEMLIST            KakaoMsgType = "FL"
	MSGTYPE_FRIENDTALK_COMMERCE            KakaoMsgType = "FM"
	MSGTYPE_FRIENDTALK_CAROUSEL_COMMERCE   KakaoMsgType = "FA"
	MSGTYPE_FRIENDTALK_VIDIO_PREMIUM       KakaoMsgType = "FP"
	MSGTYPE_BRANDMESSAGE                   KakaoMsgType = "FT"
	MSGTYPE_BRANDMESSAGE_IMAGE             KakaoMsgType = "FI"
	MSGTYPE_BRANDMESSAGE_WIDE_IMAGE        KakaoMsgType = "FW"
	MSGTYPE_BRANDMESSAGE_CAROUSEL          KakaoMsgType = "FC"
	MSGTYPE_BRANDMESSAGE_ITEMLIST          KakaoMsgType = "FL"
	MSGTYPE_BRANDMESSAGE_COMMERCE          KakaoMsgType = "FM"
	MSGTYPE_BRANDMESSAGE_CAROUSEL_COMMERCE KakaoMsgType = "FA"
	MSGTYPE_BRANDMESSAGE_VIDIO_PREMIUM     KakaoMsgType = "FP"
)

type KakaoButtonType string

const (
	KAKAO_BUTTON_WEB_LINK          KakaoButtonType = "WL"
	KAKAO_BUTTON_APP_LINK          KakaoButtonType = "AL"
	KAKAO_BUTTON_BOT_KEYWORD       KakaoButtonType = "BK"
	KAKAO_BUTTON_MESSAGE_DELIVERY  KakaoButtonType = "MD"
	KAKAO_BUTTON_DELIVERY_SCAN     KakaoButtonType = "DS"
	KAKAO_BUTTON_CONSULTBOT_CHANGE KakaoButtonType = "BC"
	KAKAO_BUTTON_CHATBOT_CHANGE    KakaoButtonType = "BT"
	KAKAO_BUTTON_ADD_CHANNEL       KakaoButtonType = "AC"
	KAKAO_BUTTON_BIZ_FORM          KakaoButtonType = "BF"
)

type RcsButtonType string

const (
	ButtonURL      RcsButtonType = "URL"
	ButtonMapLoc   RcsButtonType = "MAP_LOC"
	ButtonMapQry   RcsButtonType = "MAP_QRY"
	ButtonMapSend  RcsButtonType = "MAP_SEND"
	ButtonCalendar RcsButtonType = "CALENDAR"
	ButtonCopy     RcsButtonType = "COPY"
	ButtonCOMT     RcsButtonType = "COM_T"
	ButtonCOMV     RcsButtonType = "COM_V"
	ButtonDial     RcsButtonType = "DIAL"
)

type KakaoSendType string

const (
	BRANDMESSAGE_SENDTYPE_BASIC KakaoSendType = "basic"
	BRANDMESSAGE_SENDTYPE_FREE  KakaoSendType = "free"
)

type SimpleSendResponse struct {
	Code   string `json:"code"`
	Result string `json:"result"`
	Ref    string `json:"ref,omitempty"`
	MsgKey string `json:"msgKey,omitempty"`
}

type OmniSendResponse struct {
	Code   string               `json:"code"`
	Result string               `json:"result"`
	Ref    string               `json:"ref,omitempty"`
	Data   OmniSendResponseData `json:"data,omitempty"`
}

type OmniSendResponseData struct {
	Destinations []DestinationsResponse `json:"destinations,omitempty"`
}

type DestinationsResponse struct {
	To     string `json:"to"`
	MsgKey string `json:"msgKey"`
	Code   string `json:"code"`
	Result string `json:"result"`
}

type ReportResponse struct {
	Code   string             `json:"code"`
	Result string             `json:"result"`
	Data   ReportResponseData `json:"data,omitempty"`
}

type ReportResponseData struct {
	ReportId string   `json:"reportId,omitempty"`
	Report   []Report `json:"report"`
}

type Report struct {
	MsgKey      string `json:"msgKey"`
	ServiceType string `json:"serviceType"`
	MsgType     string `json:"msgType"`
	ReportType  string `json:"reportType"`
	ReportCode  string `json:"reportCode"`
	ReportText  string `json:"reportText,omitempty"`
	ReportTime  string `json:"reportTime"`
	SendTime    string `json:"sendTime"`
	UserType    string `json:"userType"`
	Carrier     string `json:"carrier,omitempty"`
	Ref         string `json:"ref,omitempty"`
	ResCnt      string `json:"resCnt,omitempty"`
}

type FileResponse struct {
	Code   string           `json:"code"`
	Result string           `json:"result"`
	Data   FileResponseData `json:"data,omitempty"`
}

type FileResponseData struct {
	ImgUrl  string `json:"imgUrl"`
	FileKey string `json:"fileKey"`
	Media   string `json:"media"`
	Expired string `json:"expired"`
}

type FormResponse struct {
	Code   string           `json:"code"`
	Result string           `json:"result"`
	Data   FormResponseData `json:"data,omitempty"`
}

type FormResponseData struct {
	FormId      string        `json:"formId,omitempty"`
	Expired     string        `json:"expired,omitempty"`
	MessageForm []OmniMessage `json:"messageForm"`
}

type MessageForm struct {
	MessageForm []OmniMessage `json:"messageForm"`
}
