package main

import (
	"flag"
	"fmt"
	"log"
	"zero-chat/chat/api/internal/common/imserver"
	"zero-chat/chat/api/internal/config"
	"zero-chat/chat/api/internal/handler"
	"zero-chat/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//ct, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	//defer cancel()
	//go tmp(ct, ctx)

	imServer, err := imserver.NewImServer(ctx.Redis)
	log.Printf("imServer:s%", imServer)
	if err != nil {
		log.Fatal(err)
	}
	go imServer.Subscribe()
	go imServer.Run()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//func tmp(ctx context.Context, server *svc.ServiceContext) {
//	fmt.Println("start sub")
//	for {
//		sub := server.Redis.Subscribe(ctx, "2")
//		message, err := sub.ReceiveMessage(ctx)
//		if err != nil {
//			fmt.Println("err main:", err)
//		}
//		fmt.Println("main received msg:", message)
//	}
//}
