package tunnel

import (
	"github.com/lingyicute/yiclashcore/tunnel"
)

func QueryMode() string {
	return tunnel.Mode().String()
}
