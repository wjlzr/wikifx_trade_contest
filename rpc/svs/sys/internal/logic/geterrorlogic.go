package logic

import (
	"context"

	"wikifx-trade-contest/rpc/svs/sys/internal/svc"
	"wikifx-trade-contest/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetErrorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetErrorLogic {
	return &GetErrorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  GetError todo 临时占用后期改进
func (l *GetErrorLogic) GetError(in *sys.I18NErrorReq) (*sys.I18NErrorResp, error) {

	result := l.svcCtx.ErrorModel.GetError(in.Code, in.Lang)

	return &sys.I18NErrorResp{Result: result}, nil
}
