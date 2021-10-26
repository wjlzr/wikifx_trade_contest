package oauthtypes

type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id           int64  `json:"id"`
	UserName     string `json:"userName"`
	AccessToken  string `json:"token"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

// 发送验证码
type SendCodeReq struct {
	Code            string `json:"code"`
	Phone           string `json:"phone"`
	LanguageCode    string `json:"languageCode"`
	SmsBusinessType int64  `json:"smsBusinessType"`
}

type SendCodeRes struct {
	Success   bool   `json:"success"`
	Requestid string `json:"requestid"`
}
