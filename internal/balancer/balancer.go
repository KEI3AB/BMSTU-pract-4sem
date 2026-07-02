package balancer

import "go-load-balancer/internal/pool"

type Balancer interface {
	Next() *pool.Backend
}
