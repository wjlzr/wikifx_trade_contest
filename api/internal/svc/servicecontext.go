package svc

import (
	"wikifx-trade-contest/api/internal/config"
	"wikifx-trade-contest/api/internal/middleware"
	"wikifx-trade-contest/rpc/svs/oauth/oauthclient"
	"wikifx-trade-contest/rpc/svs/sys/sysclient"

	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	Customintercept rest.Middleware
	Redis           *redis.Redis
	Oauth           oauthclient.Oauth
	Sys             sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {

	newRedis := redis.NewRedis(c.Redis.Address, redis.NodeType)

	return &ServiceContext{
		Config:          c,
		Customintercept: middleware.NewCustominterceptMiddleware().Handle,
		Redis:           newRedis,
		Oauth:           oauthclient.NewOauth(zrpc.MustNewClient(c.OauthRpc)),
		Sys:             sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
