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
