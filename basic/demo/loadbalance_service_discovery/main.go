package main

import (
	"fmt"
	"net/http"

	"github.com/yaolixiao/microservice_gateway/basic/proxy"
	"github.com/yaolixiao/microservice_gateway/basic/proxy/loadbalance"
)

var addr = "127.0.0.1:2002"

func main() {
	conf, err := loadbalance.NewLoadBalanceZkConf(
		"/rs_register",
		[]string{"127.0.0.1:2181"},
		map[string]string{},
	)

	if err != nil {
		fmt.Println("NewLoadBalanceZkConf err=", err)
		return
	}

	lb := loadbalance.LoadBalanceFactoryWithConf(loadbalance.WEIGHTROUNDROBIN, conf)
	reverseProxy := proxy.NewMultipleHostsReverseProxy(lb)
	fmt.Println("starting server at", addr)
	http.ListenAndServe(addr, reverseProxy)
}
