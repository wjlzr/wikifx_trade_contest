package logic

import (
	"context"

	"wikifx-trade-contest/rpc/svs/oauth/internal/svc"
	"wikifx-trade-contest/rpc/svs/oauth/oauth"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserCenterValidatePhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCenterValidatePhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCenterValidatePhoneLogic {
	return &UserCenterValidatePhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  验证手机号是否注册过
func (l *UserCenterValidatePhoneLogic) UserCenterValidatePhone(in *oauth.UserCenterValidatePhoneReq) (*oauth.UserCenterValidatePhoneResp, error) {
	// todo: add your logic here and delete this line

	return &oauth.UserCenterValidatePhoneResp{}, nil
}
