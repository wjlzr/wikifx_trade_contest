// Code generated by goctl. DO NOT EDIT!
// Source: oauth.proto

//go:generate mockgen -destination ./oauth_mock.go -package oauthclient -source $GOFILE

package oauthclient

import (
	"context"

	"wikifx-trade-contest/rpc/svs/oauth/oauth"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	SendCodeDataResp            = oauth.SendCodeDataResp
	SendCodeDataResultResp      = oauth.SendCodeDataResultResp
	LoginReq                    = oauth.LoginReq
	LoginResp                   = oauth.LoginResp
	UserCenterValidatePhoneReq  = oauth.UserCenterValidatePhoneReq
	UserCenterValidatePhoneResp = oauth.UserCenterValidatePhoneResp
	UserCenterSendCodeReq       = oauth.UserCenterSendCodeReq
	UserCenterSendCodeResp      = oauth.UserCenterSendCodeResp

	Oauth interface {
		//  Login 登录
		Login(ctx context.Context, in *LoginReq) (*LoginResp, error)
		//  验证手机号是否注册过
		UserCenterValidatePhone(ctx context.Context, in *UserCenterValidatePhoneReq) (*UserCenterValidatePhoneResp, error)
		//  发送验证码
		UserCenterSendCode(ctx context.Context, in *UserCenterSendCodeReq) (*UserCenterSendCodeResp, error)
	}

	defaultOauth struct {
		cli zrpc.Client
	}
)

func NewOauth(cli zrpc.Client) Oauth {
	return &defaultOauth{
		cli: cli,
	}
}

//  Login 登录
func (m *defaultOauth) Login(ctx context.Context, in *LoginReq) (*LoginResp, error) {
	client := oauth.NewOauthClient(m.cli.Conn())
	return client.Login(ctx, in)
}

//  验证手机号是否注册过
func (m *defaultOauth) UserCenterValidatePhone(ctx context.Context, in *UserCenterValidatePhoneReq) (*UserCenterValidatePhoneResp, error) {
	client := oauth.NewOauthClient(m.cli.Conn())
	return client.UserCenterValidatePhone(ctx, in)
}

//  发送验证码
func (m *defaultOauth) UserCenterSendCode(ctx context.Context, in *UserCenterSendCodeReq) (*UserCenterSendCodeResp, error) {
	client := oauth.NewOauthClient(m.cli.Conn())
	return client.UserCenterSendCode(ctx, in)
}
