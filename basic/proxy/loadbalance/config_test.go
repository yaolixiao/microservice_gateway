package loadbalance

import (
	"fmt"
	"testing"
)

func TestNewLoadBalanceObserver(t *testing.T) {
	fmt.Println("111")
	lbzkconf, err := NewLoadBalanceZkConf(
		"/rs_register",
		[]string{"127.0.0.1:2181"},
		map[string]string{"127.0.0.1:2003": "20"},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	o := NewLoadBalanceObserver(lbzkconf)
	lbzkconf.Attach(o)
	lbzkconf.UpdateConf([]string{"127.11.11"})
	select {}
}
