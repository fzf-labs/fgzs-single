package main

import (
	"fgzs-single/internal/app/web/internal/config"
	"fgzs-single/internal/app/web/internal/handler"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/middleware"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/web.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	//全局中间件
	server.Use(middleware.NewRequestLogMiddleware().Handle)
	handler.RegisterHandlers(server, ctx)
	//禁用系统日志
	if !(c.Mode == service.TestMode || c.Mode == service.ProMode) {
		logx.DisableStat()
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
