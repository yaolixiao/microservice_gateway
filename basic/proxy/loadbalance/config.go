package loadbalance

import (
	"fmt"

	"github.com/yaolixiao/microservice_gateway/basic/proxy/zookeeper"
)

type LoadBalanceConf interface {
	Attach(Observer)
	GetConf() []string
	WatchConf()
	UpdateConf([]string)
}

type Observer interface {
	Update()
}

type LoadBalanceZkConf struct {
	observers    []Observer
	path         string
	zkHosts      []string
	confIpWeight map[string]string
	activeList   []string
}

func (this *LoadBalanceZkConf) Attach(o Observer) {
	this.observers = append(this.observers, o)
}

func (this *LoadBalanceZkConf) NotifyAllObservers() {
	for _, o := range this.observers {
		o.Update()
	}
}

func (this *LoadBalanceZkConf) GetConf() []string {
	confList := []string{}
	for _, ip := range this.activeList {
		w, ok := this.confIpWeight[ip]
		if !ok {
			w = "50"
		}
		confList = append(confList, fmt.Sprintf("http://%s/base", ip)+","+w)
	}
	return confList
}

func (this *LoadBalanceZkConf) WatchConf() {
	zkManager := zookeeper.NewZkManager(this.zkHosts)
	zkManager.BuildConnect()
	chanList, chanErr := zkManager.WatchServiceList(this.path)
	go func() {
		defer zkManager.Close()
		for {
			select {
			case changedErr := <-chanErr:
				fmt.Println("changedErr:", changedErr)
			case changedList := <-chanList:
				this.UpdateConf(changedList)
			}
		}
	}()
}

func (this *LoadBalanceZkConf) UpdateConf(conf []string) {
	this.activeList = conf
	for _, o := range this.observers {
		o.Update()
	}
}

func NewLoadBalanceZkConf(path string, zkHosts []string, ipw map[string]string) (*LoadBalanceZkConf, error) {
	zkManager := zookeeper.NewZkManager(zkHosts)
	zkManager.BuildConnect()
	defer zkManager.Close()
	zlist, err := zkManager.GetServiceList(path)
	if err != nil {
		return nil, err
	}
	mConf := &LoadBalanceZkConf{
		path:         path,
		zkHosts:      zkHosts,
		confIpWeight: ipw,
		activeList:   zlist,
	}
	mConf.WatchConf()
	return mConf, nil
}

type LoadBalanceObserver struct {
	obConf *LoadBalanceZkConf
}

func (this *LoadBalanceObserver) Update() {
	fmt.Println("LoadBalanceObserver Update:", this.obConf.GetConf())
}

func NewLoadBalanceObserver(conf *LoadBalanceZkConf) *LoadBalanceObserver {
	return &LoadBalanceObserver{
		obConf: conf,
	}
}
