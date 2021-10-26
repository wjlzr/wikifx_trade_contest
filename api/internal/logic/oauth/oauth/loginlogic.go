package oauth

import (
	"context"
	"encoding/json"
	"strings"
	"wikifx-trade-contest/api/internal/svc"
	"wikifx-trade-contest/api/internal/types"
	"wikifx-trade-contest/api/internal/types/oauthtypes"
	"wikifx-trade-contest/common/errorx"
	"wikifx-trade-contest/common/response"
	"wikifx-trade-contest/rpc/svs/oauth/oauthclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据用户名和密码登录
func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {

	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, response.Error(errorx.ERROR_USERNAME_PASSWORD_NOT_EMPTY)
	}

	resp, err := l.svcCtx.Oauth.Login(l.ctx, &oauthclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, err.Error())
		return nil, response.Error(errorx.ERROR_FIND_USER_EXEPTION)
	}

	return &types.LoginResp{
		Id:           resp.Id,
		UserName:     resp.UserName,
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}, nil
}

// 发送验证码（普通）
func (l *LoginLogic) SendCode(req oauthtypes.SendCodeReq) (*oauthtypes.SendCodeRes, error) {

	if req.Code == "" || req.Phone == "" || req.LanguageCode == "" || req.SmsBusinessType == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("请求参数错误,参数:%s", reqStr)
		return nil, response.Error(errorx.ERROR_USERNAME_PASSWORD_NOT_EMPTY)
	}

	resp, err := l.svcCtx.Oauth.UserCenterSendCode(l.ctx, &oauthclient.UserCenterSendCodeReq{
		AreaCode:        req.Code,
		Phone:           req.Phone,
		LanguageCode:    req.LanguageCode,
		SmsBusinessType: req.SmsBusinessType,
	})

	if err != nil || resp.Data.Succeed != true {
		if resp.Data.Message == "" {
			return nil, response.Error(errorx.ERROR_FIND_USER_EXEPTION) // 发送验证码失败,请重试
		} else {
			return nil, response.Error(errorx.ERROR_FIND_USER_EXEPTION)
		}
	}

	return &oauthtypes.SendCodeRes{
		Success:   true,
		Requestid: resp.Data.Result.Requestid,
	}, nil
}
