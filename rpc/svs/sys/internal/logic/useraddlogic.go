package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
	"wikifx-trade-contest/library/bcryptx"
	"wikifx-trade-contest/rpc/model/user"
	"wikifx-trade-contest/rpc/svs/sys/internal/svc"
	"wikifx-trade-contest/rpc/svs/sys/sys"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserAdd 新增用户
func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (userAddResp *sys.UserAddResp, err error) {
	var users user.User
	if err = copier.Copy(&users, &in); err != nil {
		logx.WithContext(l.ctx).Errorf("svs sys logic userAdd copy Err")
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	users.Password = "123456"
	users.Salt = bcryptx.HashAndSalt("123456")

	result, _, err := l.svcCtx.UserModel.Create(users)
	if err != nil {
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return &sys.UserAddResp{
		Id:       result.Id,
		Name:     result.Name,
		NickName: result.NickName,
		IsAdmin:  result.IsAdmin,
	}, nil
}
