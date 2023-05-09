package main

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kalandramo/appdemo/cmd/user/infras/mysql"
	"github.com/kalandramo/appdemo/kitex_gen/user/userservice"
	"github.com/kalandramo/appdemo/pkg/constants"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gorm.io/plugin/opentelemetry/provider"
)

func Init() {
	mysql.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr(consts.TCP, consts.UserServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.UserServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)

	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
