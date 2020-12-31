package main

import (
	"flag"
	"fmt"

	"github.com/bjhuang103/go_pay/union_gateway/internal/config"
	"github.com/bjhuang103/go_pay/union_gateway/internal/handler"
	"github.com/bjhuang103/go_pay/union_gateway/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/bookstore-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
