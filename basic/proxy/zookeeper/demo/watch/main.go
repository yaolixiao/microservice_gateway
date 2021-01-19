package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yaolixiao/microservice_gateway/basic/proxy/zookeeper"
)

func main() {
	zkManager := zookeeper.NewZkManager([]string{"127.0.0.1:2181"})
	zkManager.BuildConnect()
	defer zkManager.Close()

	// list, err := zkManager.GetServiceList("/rs_register")
	// if err != nil {
	// 	fmt.Println("GetServiceList:", err)
	// 	return
	// }
	// fmt.Println(list)

	// 动态监听节点变化
	chanlist, chanerr := zkManager.WatchServiceList("/rs_register")
	go func() {
		for {
			select {
			case changeerr := <-chanerr:
				fmt.Println("changeerr=", changeerr)
			case changeList := <-chanlist:
				fmt.Println("watch node changed", changeList)
			}
		}
	}()

	// data, _, err := zkManager.Get("/real_node")
	// if err != nil {
	// 	fmt.Println("Get:", err)
	// 	return
	// }
	// fmt.Println(data)

	// chanlist, chanerr := zkManager.WatchPathData("/real_node")
	// go func() {
	// 	for {
	// 		select {
	// 		case changeerr := <-chanerr:
	// 			fmt.Println("changeerr=", changeerr)
	// 		case nodeContent := <-chanlist:
	// 			fmt.Println("watch content changed", string(nodeContent))
	// 		}
	// 	}
	// }()

	//关闭信号监听
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
