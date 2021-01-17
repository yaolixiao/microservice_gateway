package loadbalance

import (
	"errors"
)

// 轮询负载均衡
// 2021-01-05

type RoundRobinBalance struct {
	curIndex int
	rss      []string

	// 观察者模式
	// conf LoadBalanceConf
}

func (this *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least.")
	}

	this.rss = append(this.rss, params[0])
	return nil
}

func (this *RoundRobinBalance) Next() string {
	if len(this.rss) == 0 {
		return ""
	}

	lens := len(this.rss)
	if this.curIndex >= lens {
		this.curIndex = 0
	}
	addr := this.rss[this.curIndex]
	this.curIndex = (this.curIndex + 1) % lens
	return addr
}

// func main() {
// 	//
// 	rrb := &RoundRobinBalance{}
// 	rrb.Add("127.0.0.1:8001")
// 	rrb.Add("127.0.0.1:8002")
// 	rrb.Add("127.0.0.1:8003")

// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// 	fmt.Println(rrb.Next())
// }
