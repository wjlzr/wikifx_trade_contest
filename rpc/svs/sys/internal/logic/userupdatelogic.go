package logic

import (
	"context"

	"wikifx-trade-contest/rpc/svs/sys/internal/svc"
	"wikifx-trade-contest/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UserUpdateResp{}, nil
}
