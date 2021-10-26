package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
	"wikifx-trade-contest/api/internal/svc"
	"wikifx-trade-contest/api/internal/types"
	"wikifx-trade-contest/common/errorx"
	"wikifx-trade-contest/common/response"
	"wikifx-trade-contest/rpc/svs/sys/sysclient"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserAddLogic {
	return UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserAdd 新增用户
func (l *UserAddLogic) UserAdd(req types.AddUserReq) (*types.AddUserResp, error) {

	_, err := l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Mobile:   req.Mobile,
		Name:     req.Name,
		NickName: req.NickName,
		IsAdmin:  req.IsAdmin,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, response.Error(errorx.ERROR_USER_CREATE_FAIL)
	}

	return &types.AddUserResp{
		Code:    http.StatusOK,
		Message: "添加用户成功",
	}, nil
}
