package pool

import (
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL         *url.URL
	alive       bool
	mux         sync.RWMutex
	connections uint64
}

func NewBackend(u *url.URL) *Backend {
	return &Backend{
		URL:   u,
		alive: true,
	}
}

func (b *Backend) IncConnections() {
	atomic.AddUint64(&b.connections, 1)
}

func (b *Backend) DecConnections() {
	atomic.AddUint64(&b.connections, ^uint64(0))
}

func (b *Backend) GetActiveConnections() uint64 {
	return atomic.LoadUint64(&b.connections)
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.alive = alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	alive := b.alive
	b.mux.RUnlock()
	return alive
}
