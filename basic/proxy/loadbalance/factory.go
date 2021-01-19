package loadbalance

type LbType int

const (
	RANDOM LbType = iota
	ROUNDROBIN
	WEIGHTROUNDROBIN
)

type LoadBalance interface {
	Add(...string) error
	Next() string
}

func LoadBalanceFactory(lbType LbType) LoadBalance {

	switch lbType {
	case RANDOM:
		return &RandomBalance{}
	case ROUNDROBIN:
		return &RoundRobinBalance{}
	case WEIGHTROUNDROBIN:
		return &WeightRoundRobinBalance{}
	default:
		return &RandomBalance{}
	}
}

func LoadBalanceFactoryWithConf(lbType LbType, conf LoadBalanceConf) LoadBalance {
	if lbType == WEIGHTROUNDROBIN {
		lb := &WeightRoundRobinBalance{}
		lb.SetConf(conf)
		conf.Attach(lb)
		lb.Update()
		return lb
	}
	return LoadBalanceFactory(lbType)
}
