package models

type AlimTalkAttachment struct {
	Button        []KakaoButton  `json:"button,omitempty"`
	Item          *KakaoItemInfo `json:"item,omitempty"`
	ItemHighlight *KakaoItem     `json:"itemHighlight,omitempty"`
}

type FriendTalkAttachment struct {
	Button   []KakaoButton  `json:"button,omitempty"`
	Image    *KakaoImage    `json:"image,omitempty"`
	Item     *KakaoItemInfo `json:"item,omitempty"`
	Coupon   *KakaoCoupon   `json:"coupon,omitempty"`
	Commerce *KakaoCommerce `json:"commerce,omitempty"`
	Video    *KakaoVideo    `json:"video,omitempty"`
}

type KakaoCoupon struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	UrlPC         string `json:"url_pc,omitempty"`
	UrlMobile     string `json:"urlMobile,omitempty"`
	SchemeIOS     string `json:"schemeIOS,omitempty"`
	SchemeAndroid string `json:"schemeAndroid,omitempty"`
}

type KakaoCommerce struct {
	Title         string `json:"title"`
	RegularPrice  int    `json:"regular_price"`
	DiscountPrice int    `json:"discount_price,omitempty"`
	DiscountRate  int    `json:"discount_rate,omitempty"`
	DiscountFixed int    `json:"discount_fixed,omitempty"`
}

type KakaoVideo struct {
	VideoUrl     string `json:"video_url"`
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
}

type KakaoSuplement struct {
	QuickReply []KakaoButton `json:"quickReply"`
}

type KakaoButton struct {
	Type          KakaoButtonType `json:"type" validate:"required"`
	Name          string          `json:"name" validate:"required"`
	UrlPC         string          `json:"urlPC,omitempty"`
	UrlMobile     string          `json:"urlMobile,omitempty"`
	SchemeIOS     string          `json:"schemeIOS,omitempty"`
	SchemeAndroid string          `json:"schemeAndroid,omitempty"`
	Target        string          `json:"target,omitempty"`
	ChatExtra     string          `json:"chatExtra,omitempty"`
	ChatEvent     string          `json:"chatEvent,omitempty"`
	BizFormKey    string          `json:"bizFormKey,omitempty"`
	BizFormId     string          `json:"bizFormId,omitempty"`
}

type KakaoItemInfo struct {
	List    []KakaoItem `json:"list,omitempty"`
	Summary KakaoItem   `json:"summary,omitempty"`
}

type KakaoItem struct {
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	ImgUrl        string `json:"imgUrl,omitempty"`
	UrlPc         string `json:"url_pc,omitempty"`
	UrlMobile     string `json:"url_mobile"`
	SchemeIos     string `json:"scheme_ios,omitempty"`
	SchemeAndroid string `json:"scheme_android,omitempty"`
}

type KakaoImage struct {
	ImgUrl  string `json:"imgUrl" validate:"required"`
	ImgLink string `json:"imgLink,omitempty"`
}

type KakaoCarouselHead struct {
	Header        string `json:"header,omitempty"`
	Content       string `json:"content,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	UrlPc         string `json:"url_pc,omitempty"`
	UrlMobile     string `json:"url_mobile"`
	SchemeIos     string `json:"scheme_ios,omitempty"`
	SchemeAndroid string `json:"scheme_android,omitempty"`
}

type KakaoCarouselList struct {
	Header            string                `json:"header,omitempty"`
	Message           string                `json:"message,omitempty"`
	AdditionalContent string                `json:"additional_content,omitempty"`
	Attachment        *FriendTalkAttachment `json:"attachment,omitempty"`
}

type KakaoCarouselTail struct {
	UrlPc         string `json:"url_pc,omitempty"`
	UrlMobile     string `json:"url_mobile,omitempty"`
	SchemeIos     string `json:"scheme_ios,omitempty"`
	SchemeAndroid string `json:"scheme_android,omitempty"`
}
type KakaoCarousel struct {
	Head *KakaoCarouselHead  `json:"head,omitempty"`
	List []KakaoCarouselList `json:"list,omitempty"`
	Tail *KakaoCarouselTail  `json:"tail,omitempty"`
}
