package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	proxy_addr = "http://127.0.0.1:2003"
	port       = "2002"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// 1. 解析代理地址，更改请求体协议和主机
	proxy, _ := url.Parse(proxy_addr)
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	// 2. 请求下游
	transport := http.DefaultTransport
	res, err := transport.RoundTrip(r)
	if err != nil {
		log.Println("请求下游数据失败，", err)
		return
	}

	// 3. 将请求回来的数据返回给上游
	for k, vv := range res.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}

	defer res.Body.Close()
	bufio.NewReader(res.Body).WriteTo(w)
}

func main() {
	// 启动代理服务器
	http.HandleFunc("/", handler)
	fmt.Println("starting server at", ":"+port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
