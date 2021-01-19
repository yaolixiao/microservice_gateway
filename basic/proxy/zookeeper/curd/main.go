package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var host = []string{"127.0.0.1:2181"}

func main() {
	//
	conn, _, err := zk.Connect(host, 5*time.Second)
	if err != nil {
		fmt.Println("zk.Connect err=", err)
		return
	}

	// str, err := conn.Create("/test_tree", []byte("tree_content"), 0, zk.WorldACL(zk.PermAll))
	// if err != nil {
	// 	fmt.Println("conn.Create err=", err)
	// 	return
	// }
	// fmt.Println("create:", str)

	nodeValue, dstat, err := conn.Get("/test_tree")
	if err != nil {
		fmt.Println("conn.Get err=", err)
		return
	}
	fmt.Println("Get:", string(nodeValue), dstat.Version)

	// dstat, err = conn.Set("/test_tree", []byte("tree_content1111"), dstat.Version)
	// if err != nil {
	// 	fmt.Println("conn.Get err=", err)
	// 	return
	// }

	err = conn.Delete("/test_tree", dstat.Version)
	if err != nil {
		fmt.Println("conn.Delete err=", err)
		return
	}

	nodeValue, dstat, err = conn.Get("/test_tree")
	if err != nil {
		fmt.Println("conn.Get err=", err)
		return
	}
	fmt.Println("Get:", string(nodeValue), dstat.Version)

	hasNode, _, err := conn.Exists("/test_tree")
	if err != nil {
		fmt.Println("conn.Exists err=", err)
		return
	}
	fmt.Println("hasNode:", hasNode)
}
