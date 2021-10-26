package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"wikifx-trade-contest/rpc/svs/oauth/internal/client"

	"github.com/k0kubun/pp"

	"wikifx-trade-contest/rpc/svs/oauth/internal/svc"
	"wikifx-trade-contest/rpc/svs/oauth/oauth"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserCenterSendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCenterSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCenterSendCodeLogic {
	return &UserCenterSendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  发送验证码
func (l *UserCenterSendCodeLogic) UserCenterSendCode(in *oauth.UserCenterSendCodeReq) (res *oauth.UserCenterSendCodeResp, err error) {

	jsonStr, _ := json.Marshal(client.SendCodeRequest{AreaCode: in.AreaCode, Phone: in.Phone, LanguageCode: in.LanguageCode, UserId: in.UserId, SmsBusinessType: in.SmsBusinessType, ApplicationType: client.ApplicationType})
	request, err := client.Request(http.MethodPost, l.svcCtx.Config.UserCenter.User+client.SendCode, bytes.NewBuffer(jsonStr), l.svcCtx.Config)
	if err != nil {
		logx.Errorf("UserCenter SendCode 请求 err：%s", err.Error())
		return &oauth.UserCenterSendCodeResp{}, err
	}

	content := client.ResponseHandle(request)
	pp.Println(string(content))
	_ = json.Unmarshal(content, &res)

	if res.Code != 200 || res.Success != true {
		logx.Info("UserCenter SendCode 发送验证码Error response：", res)
		return &oauth.UserCenterSendCodeResp{}, errors.New(res.Msg)
	}
	pp.Println(res)
	return res, nil
}
