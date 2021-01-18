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

	"github.com/yaolixiao/microservice_gateway/basic/proxy/https-server/tdata"

	"golang.org/x/net/http2"
)

// CA私钥
// openssl genrsa -out ca.key 2048
// CA数据证书
// openssl req -x509 -new -nodes -key ca.key -subj "/CN=example1.com" -days 5000 -out ca.crt
// 服务器私钥（默认由CA签发）
// openssl genrsa -out server.key 2048
// 服务器证书签名请求 certificate sign request 简称csr
// openssl req -new -key server.key -subj "/CN=example1.com" -out server.csr
// 上面2个文件生成服务器证书（days代表有效期）
// openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

type RealServer struct {
	Addr string
}

func (this *RealServer) Run() {
	log.Println("starting server at ", this.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", this.HelloHandler)
	mux.HandleFunc("/base/error", this.ErrorHandler)
	server := &http.Server{
		Addr:         this.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() {
		// 将http升级为http2
		http2.ConfigureServer(server, &http2.Server{})
		log.Fatal(server.ListenAndServeTLS(tdata.Path("server.crt"), tdata.Path("server.key")))
	}()
}

func (this *RealServer) HelloHandler(rw http.ResponseWriter, req *http.Request) {
	upath := fmt.Sprintf("http://%s%s\n", this.Addr, req.URL.Path)
	io.WriteString(rw, upath)
}

func (this *RealServer) ErrorHandler(rw http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	rw.WriteHeader(500)
	io.WriteString(rw, upath)
}

func main() {
	//
	rs1 := &RealServer{Addr: "127.0.0.1:3003"}
	rs1.Run()
	rs2 := &RealServer{Addr: "127.0.0.1:3004"}
	rs2.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
