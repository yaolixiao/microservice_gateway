package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yaolixiao/microservice_gateway/basic/proxy/zookeeper"
)

// 网关拓展服务发现
// 1. 下游机器启动时创建临时节点：节点名与内容为服务地址
// 2. 以观察者模式构建负载均衡配置 LoadBalanceConf
// 3. LoadBalanceConf 与负载均衡器整合

// 给代理提供下游服务器
type realserver struct {
	addr string
}

// 启动方法
func (rs *realserver) run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rs.helloHandler)
	mux.HandleFunc("/base/error", rs.errorHandler)

	server := &http.Server{
		Addr:         rs.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}

	fmt.Println("starting server at", rs.addr)
	go func() {
		// 注册zk，创建临时节点
		zkManager := zookeeper.NewZkManager([]string{"127.0.0.1:2181"})
		err := zkManager.BuildConnect()
		if err != nil {
			fmt.Println("BuildConnect err=", err)
			return
		}
		defer zkManager.Close()
		err = zkManager.CreateTempNode("/rs_register", rs.addr)
		if err != nil {
			fmt.Println("CreateTempNode err=", err)
			return
		}
		zlist, err := zkManager.GetServiceList("/rs_register")
		if err != nil {
			fmt.Println("GetServiceList err=", err)
			return
		}
		fmt.Println(zlist)

		// 启动服务
		log.Fatal(server.ListenAndServe())
	}()
}

func (rs *realserver) helloHandler(w http.ResponseWriter, r *http.Request) {
	upath := fmt.Sprintf("hello handler http://%s%s\n", rs.addr, r.URL.Path)
	io.WriteString(w, upath)
}

func (rs *realserver) errorHandler(w http.ResponseWriter, r *http.Request) {
	upath := fmt.Sprintf("error handler http://%s%s\n", rs.addr, r.URL.Path)
	w.WriteHeader(500)
	io.WriteString(w, upath)
}

func main() {
	// 创建服务器
	rs1 := &realserver{addr: "127.0.0.1:2003"}
	rs1.run()
	time.Sleep(time.Second * 20)
	rs2 := &realserver{addr: "127.0.0.1:2004"}
	rs2.run()

	// 监听关闭信号，syscall.SIGINT, syscall.SIGTERM
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
