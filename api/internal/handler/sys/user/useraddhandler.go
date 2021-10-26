package handler

import (
	"github.com/k0kubun/pp"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	logic "wikifx-trade-contest/api/internal/logic/sys/user"
	"wikifx-trade-contest/api/internal/svc"
	"wikifx-trade-contest/api/internal/types"
	"wikifx-trade-contest/common/response"
)

func UserAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			pp.Println(err)
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserAddLogic(r.Context(), ctx)
		resp, err := l.UserAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
