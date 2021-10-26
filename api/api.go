package main

import (
	"flag"
	"fmt"
	"net/http"
	"wikifx-trade-contest/api/internal/config"
	"wikifx-trade-contest/api/internal/handler"
	"wikifx-trade-contest/api/internal/svc"
	"wikifx-trade-contest/common/response"

	"github.com/tal-tech/go-zero/rest/httpx"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var (
	configFile = flag.String("f", "api/etc/api.yaml", "the config file")
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(
		func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.Error(w, response.Error(http.StatusUnauthorized))
		}))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *response.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
