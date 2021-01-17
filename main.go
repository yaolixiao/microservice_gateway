package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yaolixiao/microservice_gateway/dao"

	"github.com/yaolixiao/microservice_gateway/http_proxy_router"

	"github.com/yaolixiao/golang_common/lib"
	"github.com/yaolixiao/microservice_gateway/router"
)

var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	configs  = flag.String("configs", "", "input config file like ./conf/dev/")
)

func main() {

	// 取出命令行参数
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *configs == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *endpoint == "dashboard" {
		if err := lib.InitModule(*configs, []string{"base", "redis", "mysql"}); err != nil {
			fmt.Printf("main init fail. err=%v\n", err)
			return
		}
		defer lib.Destroy()
		// 启动服务
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		// 优雅关停服务器
		// todo...
	} else if *endpoint == "server" {
		if err := lib.InitModule(*configs, []string{"base", "redis", "mysql"}); err != nil {
			fmt.Printf("main init fail. err=%v\n", err)
			return
		}
		defer lib.Destroy()
		dao.ServiceManagerHandler.LoadOnce()

		// 为了同时启动多个服务。需在协程中启动服务
		go func() { http_proxy_router.HttpServerRun() }()
		// go func() { http_proxy_router.HttpsServerRun() }()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}
}
