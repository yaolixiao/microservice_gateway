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
)

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
	go func() { log.Fatal(server.ListenAndServe()) }()
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
	rs2 := &realserver{addr: "127.0.0.1:2004"}
	rs2.run()

	// 监听关闭信号，syscall.SIGINT, syscall.SIGTERM
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
