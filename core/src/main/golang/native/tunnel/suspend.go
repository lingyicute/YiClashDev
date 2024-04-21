package tunnel

import "github.com/lingyicute/yiclashcore/adapter/provider"

func Suspend(s bool) {
	provider.Suspend(s)
}
