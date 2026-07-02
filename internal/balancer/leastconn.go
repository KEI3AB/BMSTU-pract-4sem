package balancer

import (
	"math/rand"

	"go-load-balancer/internal/pool"
)

type LeastConn struct {
	pool *pool.Pool
}

func NewLeastConn(p *pool.Pool) *LeastConn {
	return &LeastConn{
		pool: p,
	}
}

func (l *LeastConn) Next() *pool.Backend {
	backends := l.pool.GetAlive()
	if len(backends) == 0 {
		return nil
	}

	var candidates []*pool.Backend
	minConn := ^uint64(0)

	for _, b := range backends {
		c := b.GetActiveConnections()
		if c < minConn {
			minConn = c
			candidates = []*pool.Backend{b}
		} else if c == minConn {
			candidates = append(candidates, b)
		}
	}

	if len(candidates) == 1 {
		return candidates[0]
	}

	return candidates[rand.Intn(len(candidates))]
}
