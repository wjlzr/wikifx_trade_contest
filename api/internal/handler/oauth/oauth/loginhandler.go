package oauth

import (
	"net/http"
	logic "wikifx-trade-contest/api/internal/logic/oauth/oauth"
	"wikifx-trade-contest/api/internal/svc"
	"wikifx-trade-contest/api/internal/types"
	"wikifx-trade-contest/api/internal/types/oauthtypes"
	"wikifx-trade-contest/common/response"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}

// SendCodeHandler 发送验证码
func SendCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req oauthtypes.SendCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.SendCode(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
