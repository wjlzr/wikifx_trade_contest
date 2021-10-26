package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"wikifx-trade-contest/rpc/svs/oauth/internal/config"

	"github.com/tal-tech/go-zero/core/logx"
)

//init token
func passiveInit(config config.Config) error {
	resp, err := http.Get(config.UserCenter.TestUrl + fmt.Sprintf(getToken+"?username=%s&password=%s", userName, password))
	if err != nil {
		logx.Errorf("UserCenter init http err：%s", err.Error())
		return err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("UserCenter init ioutil.ReadAll err：%s", err.Error())
		return err
	}
	var t tokenResponse
	_ = json.Unmarshal(bs, &t)
	if t.Status == false || t.AccessToken == "" {
		logx.Errorf("UserCenter init 授权失败")
		return errors.New("授权失败")
	}

	authorization = t.TokenType + " " + t.AccessToken
	return nil
}

// http请求
func Request(method, url string, body io.Reader, config config.Config) (request *http.Request, err error) {

	// 签名流程
	if resp := passiveInit(config); resp != nil {
		return nil, errors.New("签名流程失败")
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logx.Errorf("UserCenter request http err：%s", err.Error())
		return request, err
	}
	req.Header.Add("Authorization", authorization)
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

//返回参数统一处理
func ResponseHandle(request *http.Request) []byte {
	client := &http.Client{}

	resp, _ := client.Do(request)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("调用中台接口：%+v \n", resp.Request)
	fmt.Printf("用户中台返回值：%s \n", string(content))
	return content
}
