package main

import (
	"fmt"
	"errors"
	"math/rand"
)

// 随机负载均衡
// 2020-01-05

type RandomBalance struct {
	curIndex int
	rss []string

	// 观察者模式
	// conf LoadBalanceConf
}

func (this *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least.")
	}

	this.rss = append(this.rss, params[0])
	return nil
}

func (this *RandomBalance) Next() string {
	if len(this.rss) == 0 {
		return ""
	}

	this.curIndex = rand.Intn(len(this.rss))
	return this.rss[this.curIndex]
}

func main() {
	rb := &RandomBalance{}
	rb.Add("127.0.0.1:8001")
	rb.Add("127.0.0.1:8002")
	rb.Add("127.0.0.1:8003")
	rb.Add("127.0.0.1:8004")
	rb.Add("127.0.0.1:8005")

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}