package dingtalk_test

import (
	"github.com/smartwalle/dingtalk"
	"testing"
)

var client *dingtalk.Client

func init() {
	client = dingtalk.New("ding0umf5qncge8e4qjv", "hello")
}

func TestClient_GetToken(t *testing.T) {
	rsp, err := client.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp.ErrCode, rsp.ErrMsg, rsp.AccessToken)
}
