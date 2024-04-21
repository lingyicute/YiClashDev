package tunnel

import (
	C "github.com/lingyicute/yiclashcore/constant"
	"github.com/lingyicute/yiclashcore/tunnel/statistic"
)

func CloseAllConnections() {
	statistic.DefaultManager.Range(func(c statistic.Tracker) bool {
		_ = c.Close()
		return true
	})
}

func closeMatch(filter func(conn C.Conn) bool) {
	statistic.DefaultManager.Range(func(c statistic.Tracker) bool {
		if cc, ok := c.(C.Conn); ok {
			if filter(cc) {
				_ = cc.Close()
				return true
			}
		}
		return false
	})
}

func closeConnByGroup(name string) {
	closeMatch(func(conn C.Conn) bool {
		for _, c := range conn.Chains() {
			if c == name {
				return true
			}
		}

		return false
	})
}
