package main

import (
	"flag"
	"fmt"

	"wikifx-trade-contest/rpc/svs/sys/internal/config"
	"wikifx-trade-contest/rpc/svs/sys/internal/server"
	"wikifx-trade-contest/rpc/svs/sys/internal/svc"
	"wikifx-trade-contest/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewSysServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sys.RegisterSysServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting sys.rpc server at %s...\n", c.ListenOn)
	s.Start()
}
