package main

import (
	"fgzs-single/internal/app/admin/internal/config"
	"fgzs-single/internal/app/admin/internal/handler"
	"fgzs-single/internal/app/admin/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)
	//禁用系统日志
	if !(c.Mode == service.TestMode || c.Mode == service.ProMode) {
		logx.DisableStat()
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
