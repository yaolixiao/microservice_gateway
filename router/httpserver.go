package router

import (
	"net/http"
	"time"
	"github.com/yaolixiao/golang_common/lib"
	"github.com/gin-gonic/gin"
	"log"
)

var HttpSrvHandle *http.Server

func HttpServerRun() {
	gin.SetMode(lib.ConfBase.DebugMode)
	router := InitRouter()

	HttpSrvHandle = &http.Server{
		Addr:			lib.GetStringConf("base.http.addr"),
		Handler:		router,
		ReadTimeout:	time.Duration(lib.GetIntConf("base.http.read_timeout")) * time.Second,
		WriteTimeout:	time.Duration(lib.GetIntConf("base.http.write_timeout")) * time.Second,
		MaxHeaderBytes:	1 << uint(lib.GetIntConf("base.http.max_header_bytes")),
	}

	go func() {
		log.Printf("[INFO] HttpServerRun: %s\n", lib.GetStringConf("base.http.addr"))
		if err := HttpSrvHandle.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] HttpServerRun listen: %s\n", err)
		}
	} ()
}