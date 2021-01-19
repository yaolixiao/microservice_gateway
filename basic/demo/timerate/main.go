package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	l := rate.NewLimiter(1, 5)
	log.Println(l.Limit(), l.Burst())
	for i := 1; i < 100; i++ {

		// 阻塞等待直到取到第一个token
		log.Println("before wait...", i)
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait error=", err)
		}
		log.Println("after wait")

		// 返回需要等待多久才有新的token，这样就可以等待指定时间执行任务
		r := l.Reserve()
		log.Println("reserve delay:", r.Delay())

		// 判断当前是否可以取到token
		a := l.Allow()
		log.Println("allow:", a)

		// time.Sleep(time.Second * 2)
		fmt.Println()
		fmt.Println()
	}
}
