package logic

import (
	"context"

	"wikifx-trade-contest/rpc/svs/sys/internal/svc"
	"wikifx-trade-contest/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *sys.UserDeleteReq) (*sys.UserDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UserDeleteResp{}, nil
}
