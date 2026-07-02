package balancer

import (
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	balancer Balancer
}

func NewProxy(b Balancer) *Proxy {
	return &Proxy{
		balancer: b,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := p.balancer.Next()
	if backend == nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	backend.IncConnections()
	defer backend.DecConnections()

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}

	proxy.ServeHTTP(w, r)
}
