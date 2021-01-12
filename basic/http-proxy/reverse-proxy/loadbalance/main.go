package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/yaolixiao/microservice_gateway/basic/proxy/loadbalance"
)

var (
	addr      = "127.0.0.1:2002"
	transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue状态码超时时间
	}
)

func NewMultipleHostsReverseProxy(lb loadbalance.LoadBalance) *httputil.ReverseProxy {

	director := func(req *http.Request) {

		url1 := lb.Next()
		if url1 == "" {
			log.Fatal("找不到下游服务器")
		}

		target, err := url.Parse(url1)
		if err != nil {
			log.Fatal("url.Parse error, err=", err)
		}

		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	// 修改返回内容
	modifyFunc := func(resp *http.Response) error {
		//
		if resp.StatusCode != 200 {
			// 获取内容
			oldBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// 追加内容
			newBody := []byte("StatusCode Error: " + string(oldBody))
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
			resp.ContentLength = int64(len(newBody))
			resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(newBody)), 10))
		}
		return nil
	}

	return &httputil.ReverseProxy{
		Director:       director,
		Transport:      transport,
		ModifyResponse: modifyFunc,
	}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func main() {
	//
	lb := loadbalance.LoadBalanceFactory(loadbalance.WEIGHTROUNDROBIN)
	lb.Add("http://127.0.0.1:2003/base", "4")
	lb.Add("http://127.0.0.1:2004/base", "2")

	proxy := NewMultipleHostsReverseProxy(lb)
	fmt.Println("starting server at", addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
