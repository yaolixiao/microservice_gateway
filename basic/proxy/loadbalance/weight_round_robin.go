package loadbalance

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 加权负载均衡
// 2021-01-05
// 权重核心计算步骤 (参考Nginx加权负载均衡)
// 1. currentWeight = currentWeight + effectiveWeight
// 2. 选中最大的currentWeight为当前节点
// 3. currentWeight = currentWeight - totalWeight(sum(effectiveWeight))

type WeightNode struct {
	addr            string
	weight          int // 权重值
	currentWeight   int // 节点当前权重
	effectiveWeight int // 有效权重
}

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int

	// 观察者模式
	conf LoadBalanceConf
}

func (this *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params len must 2")
	}

	w, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}

	node := &WeightNode{addr: params[0], weight: int(w)}
	node.effectiveWeight = node.weight
	this.rss = append(this.rss, node)
	return nil
}

func (this *WeightRoundRobinBalance) Next() string {
	var total = 0
	var best *WeightNode
	for _, node := range this.rss {
		// 1. 计算有效权重之和total
		total += node.effectiveWeight

		// 2. 计算临时节点权重，currentWeight + effectiveWeight
		node.currentWeight += node.effectiveWeight

		// 3. 有效权重默认与权重相等，节点失败-1，节点成功+1，直到等于权重weight
		if node.effectiveWeight < node.weight {
			node.effectiveWeight++
		}

		// 4. 选择最大临时权重节点
		if best == nil || node.currentWeight > best.currentWeight {
			best = node
		}
	}
	if best == nil {
		return ""
	}
	// 5. 变更临时权重为currentWeight-total
	best.currentWeight -= total
	return best.addr
}

func (this *WeightRoundRobinBalance) SetConf(conf LoadBalanceConf) {
	this.conf = conf
}

func (this *WeightRoundRobinBalance) Update() {
	if conf, ok := this.conf.(*LoadBalanceZkConf); ok {
		olist := conf.GetConf()
		fmt.Println("权重轮询观察者列表：", olist)
		for _, ipw := range olist {
			ipws := strings.Split(ipw, ",")
			this.Add(ipws...)
		}

	} else {
		fmt.Println("权重轮询观察者 类型断言失败")
	}
}

// func main() {
// 	rb := &WeightRoundRobinBalance{}
// 	rb.Add("127.0.0.1:8001", "40")
// 	rb.Add("127.0.0.1:8002", "30")
// 	rb.Add("127.0.0.1:8003", "20")

// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// 	fmt.Println(rb.Next())
// }
