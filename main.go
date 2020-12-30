package main

import (
	"fmt"
	"github.com/yaolixiao/golang_common/lib"
	"github.com/yaolixiao/microservice_gateway/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := lib.InitModule("./conf/dev/", []string{"base", "redis", "mysql"}); err != nil {
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
}