package dingtalk

import "time"

const (
	kAPIDomain = "https://oapi.dingtalk.com"

	kContentType = "application/json"
)

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type GetTokenRsp struct {
	Response
	AccessToken string `json:"access_token"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresOn   int64  `json:"expires_on"`
}

func (this *Token) IsValid() bool {
	return time.Now().Unix()-this.ExpiresOn > 0
}
