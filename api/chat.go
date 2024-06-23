package main

import (
	"context"
	"flag"
	"fmt"

	"zero-chat/api/internal/config"
	"zero-chat/api/internal/handler"
	"zero-chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	ct := context.Background()

	go tmp(ct, ctx)

	server.Start()
}

func tmp(ctx context.Context, server *svc.ServiceContext) {
	fmt.Println("start sub")
	for {
		sub := server.Redis.Subscribe(ctx, "2")
		message, err := sub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("err main:", err)
		}
		fmt.Println("main msg:", message)
	}

}
