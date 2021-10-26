package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"wikifx-trade-contest/common/errorx"
	"wikifx-trade-contest/rpc/svs/oauth/internal/svc"
	"wikifx-trade-contest/rpc/svs/oauth/oauth"
	"time"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Login 登录
func (l *LoginLogic) Login(in *oauth.LoginReq) (*oauth.LoginResp, error) {

	userInfo, count, err := l.svcCtx.Oauth.FindOne(map[string]interface{}{"name": in.UserName, "password": in.Password})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		logx.WithContext(l.ctx).Errorf("账号或密码错误")
		return nil, errors.New(errorx.GetMsg(errorx.ERROR_USERNAME_PASSWORD_ERR))
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, userInfo.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &oauth.LoginResp{
		Id:           userInfo.Id,
		UserName:     userInfo.Name,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)
	return resp, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
