// Code generated by goctl. DO NOT EDIT!
// Source: oauth.proto

package server

import (
	"context"

	"wikifx-trade-contest/rpc/svs/oauth/internal/logic"
	"wikifx-trade-contest/rpc/svs/oauth/internal/svc"
	"wikifx-trade-contest/rpc/svs/oauth/oauth"
)

type OauthServer struct {
	svcCtx *svc.ServiceContext
}

func NewOauthServer(svcCtx *svc.ServiceContext) *OauthServer {
	return &OauthServer{
		svcCtx: svcCtx,
	}
}

//  Login 登录
func (s *OauthServer) Login(ctx context.Context, in *oauth.LoginReq) (*oauth.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

//  验证手机号是否注册过
func (s *OauthServer) UserCenterValidatePhone(ctx context.Context, in *oauth.UserCenterValidatePhoneReq) (*oauth.UserCenterValidatePhoneResp, error) {
	l := logic.NewUserCenterValidatePhoneLogic(ctx, s.svcCtx)
	return l.UserCenterValidatePhone(in)
}

//  发送验证码
func (s *OauthServer) UserCenterSendCode(ctx context.Context, in *oauth.UserCenterSendCodeReq) (*oauth.UserCenterSendCodeResp, error) {
	l := logic.NewUserCenterSendCodeLogic(ctx, s.svcCtx)
	return l.UserCenterSendCode(in)
}
