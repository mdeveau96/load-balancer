package algorithms

import (
	"context"
	"errors"
	"sync"

	"github.com/michaeldeveau/load-balancer/pkg/loadbalancer"
)

type RoundRobinLoadBalancer struct {
	servers []loadbalancer.Server
	current int
	mutex sync.RWMutex
}

func NewRoundRobinLoadBalancer() *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		servers: make([]loadbalancer.Server, 0),
		current: 0,
	}
}

func (rr *RoundRobinLoadBalancer) SelectServer(ctx context.Context) (loadbalancer.Server, error) {
	rr.mutex.Lock()
	defer rr.mutex.Unlock()

	if len(rr.servers) == 0 {
		return nil, errors.New("no servers available")
	}
