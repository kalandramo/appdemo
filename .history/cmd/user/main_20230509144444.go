package main

import (
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kalandramo/appdemo/cmd/user/infras/mysql"
	user "github.com/kalandramo/appdemo/kitex_gen/user/userservice"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func Init() {
	mysql.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
