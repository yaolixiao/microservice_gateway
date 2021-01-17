package http_proxy_router

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/golang_common/lib"
	"github.com/yaolixiao/microservice_gateway/middleware"
)

var (
	HttpSrvHandle *http.Server
	// HttpsSrvHandle *http.Server
)

func HttpServerRun() {
	gin.SetMode(lib.GetStringConf("proxy.base.debug_mode"))
	router := InitRouter(middleware.RecoveryMiddleware(), middleware.RequestLog())

	HttpSrvHandle = &http.Server{
		Addr:           lib.GetStringConf("proxy.http.addr"),
		Handler:        router,
		ReadTimeout:    time.Duration(lib.GetIntConf("proxy.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.GetIntConf("proxy.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.GetIntConf("proxy.http.max_header_bytes")),
	}

	go func() {
		log.Printf("[INFO] http_proxy_server http run: %s\n", lib.GetStringConf("proxy.http.addr"))
		if err := HttpSrvHandle.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] http_proxy_server http listen: %s\n", err)
		}
	}()
}

// func HttpsServerRun() {
// 	gin.SetMode(lib.GetStringConf("proxy.base.debug_mode"))
// 	router := InitRouter()

// 	HttpsSrvHandle = &http.Server{
// 		Addr:           lib.GetStringConf("proxy.https.addr"),
// 		Handler:        router,
// 		ReadTimeout:    time.Duration(lib.GetIntConf("proxy.https.read_timeout")) * time.Second,
// 		WriteTimeout:   time.Duration(lib.GetIntConf("proxy.https.write_timeout")) * time.Second,
// 		MaxHeaderBytes: 1 << uint(lib.GetIntConf("proxy.https.max_header_bytes")),
// 	}

// 	go func() {
// 		log.Printf("[INFO] http_proxy_server https run: %s\n", lib.GetStringConf("proxy.https.addr"))
// 		if err := HttpsSrvHandle.ListenAndServeTLS(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("[ERROR] http_proxy_server https listen: %s\n", err)
// 		}
// 	}()
// }
