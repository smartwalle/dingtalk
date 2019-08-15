package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Client struct {
	mu        sync.Mutex
	appKey    string
	appSecret string
	Client    *http.Client
	token     *Token
}

func New(appKey, appSecret string) (client *Client) {
	client = &Client{}
	client.appKey = appKey
	client.appSecret = appSecret
	client.Client = http.DefaultClient
	return client
}

func (this *Client) doRequest(method, path string, param url.Values, reqData interface{}, result interface{}) (err error) {
	var buf io.Reader
	if reqData != nil {
		reqBytes, _ := json.Marshal(reqData)
		buf = bytes.NewReader(reqBytes)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s?%s", kAPIDomain, path, param.Encode()), buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) doRequestWithAccessToken(method, path string, param url.Values, reqData interface{}, result interface{}) (err error) {
	if param == nil {
		param = url.Values{}
	}

	this.mu.Lock()
	if this.token == nil || this.token.IsValid() == false {
		if _, err = this.RefreshToken(); err != nil {
			defer this.mu.Unlock()
			return err
		}
	}
	defer this.mu.Unlock()

	param.Set("access_token", this.token.AccessToken)

	return this.doRequest(method, path, param, reqData, result)
}

func (this *Client) GetToken() (result *GetTokenRsp, err error) {
	var p = url.Values{}
	p.Set("appkey", this.appKey)
	p.Set("appsecret", this.appSecret)

	if err = this.doRequest(http.MethodGet, "gettoken", p, nil, &result); err != nil {
		return nil, err
	}

	if result != nil {
		var token = &Token{}
		token.AccessToken = result.AccessToken
		token.ExpiresOn = time.Now().Unix() + 7200
		this.token = token
	}

	return result, nil
}

func (this *Client) RefreshToken() (result *GetTokenRsp, err error) {
	return this.GetToken()
}
