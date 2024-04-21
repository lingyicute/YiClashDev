package tunnel

import (
	"sync"

	"github.com/lingyicute/yiclashcore/adapter"
	"github.com/lingyicute/yiclashcore/adapter/outboundgroup"
	"github.com/lingyicute/yiclashcore/constant/provider"
	"github.com/lingyicute/yiclashcore/log"
	"github.com/lingyicute/yiclashcore/tunnel"
)

func HealthCheck(name string) {
	p := tunnel.Proxies()[name]

	if p == nil {
		log.Warnln("Request health check for `%s`: not found", name)

		return
	}

	g, ok := p.(*adapter.Proxy).ProxyAdapter.(outboundgroup.ProxyGroup)
	if !ok {
		log.Warnln("Request health check for `%s`: invalid type %s", name, p.Type().String())

		return
	}

	wg := &sync.WaitGroup{}

	for _, pr := range g.Providers() {
		wg.Add(1)

		go func(provider provider.ProxyProvider) {
			provider.HealthCheck()

			wg.Done()
		}(pr)
	}

	wg.Wait()
}

func HealthCheckAll() {
	for _, g := range QueryProxyGroupNames(false) {
		go func(group string) {
			HealthCheck(group)
		}(g)
	}
}
