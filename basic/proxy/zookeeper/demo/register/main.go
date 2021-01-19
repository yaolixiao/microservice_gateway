package main

import (
	"fmt"
	"time"

	"github.com/yaolixiao/microservice_gateway/basic/proxy/zookeeper"
)

func main() {
	zkManager := zookeeper.NewZkManager([]string{"127.0.0.1:2181"})
	zkManager.BuildConnect()
	defer zkManager.Close()
	i := 0
	for {
		if err := zkManager.CreateTempNode("/real_node", fmt.Sprint(i)); err != nil {
			fmt.Println("CreateTempNode err=", err)
			return
		}

		fmt.Println("tempnode:", i)
		time.Sleep(time.Second * 5)
		i++
	}
}
