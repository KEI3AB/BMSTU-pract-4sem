package balancer

import (
	"sync/atomic"

	"go-load-balancer/internal/pool"
)

type RoundRobin struct {
	pool    *pool.Pool
	current uint64
}

func NewRoundRobin(p *pool.Pool) *RoundRobin {
	return &RoundRobin{
		pool: p,
	}
}

func (r *RoundRobin) Next() *pool.Backend {
	backends := r.pool.GetAlive()
	if len(backends) == 0 {
		return nil
	}

	next := atomic.AddUint64(&r.current, 1)
	idx := next % uint64(len(backends))

	return backends[idx]
}
