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
		conf := fmt.Sprintf("{name:" + fmt.Sprint(i) + "}")
		zkManager.Set("/real_node", []byte(conf), int32(i))
		fmt.Println("Write:", conf)
		time.Sleep(time.Second * 5)
		i++
	}
}
