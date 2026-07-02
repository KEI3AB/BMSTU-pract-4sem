package pool

type Pool struct {
	backends []*Backend
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) AddBackend(b *Backend) {
	p.backends = append(p.backends, b)
}

func (p *Pool) GetAlive() []*Backend {
	var alive []*Backend
	for _, b := range p.backends {
		if b.IsAlive() {
			alive = append(alive, b)
		}
	}
	return alive
}

func (p *Pool) GetAll() []*Backend {
	return p.backends
}
