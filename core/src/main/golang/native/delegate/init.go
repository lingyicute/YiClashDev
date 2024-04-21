package delegate

import (
	"errors"
	"syscall"

	"cfa/blob"

	"github.com/lingyicute/yiclashcore/component/process"
	"github.com/lingyicute/yiclashcore/log"

	"cfa/native/app"
	"cfa/native/platform"

	"github.com/lingyicute/yiclashcore/component/dialer"
	"github.com/lingyicute/yiclashcore/component/mmdb"
	"github.com/lingyicute/yiclashcore/constant"
)

var errBlocked = errors.New("blocked")

func Init(home, versionName string, platformVersion int) {
	mmdb.LoadFromBytes(blob.GeoipDatabase)
	constant.SetHomeDir(home)
	app.ApplyVersionName(versionName)
	app.ApplyPlatformVersion(platformVersion)

	process.DefaultPackageNameResolver = func(metadata *constant.Metadata) (string, error) {
		src, dst := metadata.RawSrcAddr, metadata.RawDstAddr

		if src == nil || dst == nil {
			return "", process.ErrInvalidNetwork
		}

		uid := app.QuerySocketUid(metadata.RawSrcAddr, metadata.RawDstAddr)
		pkg := app.QueryAppByUid(uid)

		log.Debugln("[PKG] %s --> %s by %d[%s]", metadata.SourceAddress(), metadata.RemoteAddress(), uid, pkg)

		return pkg, nil
	}

	dialer.DefaultSocketHook = func(network, address string, conn syscall.RawConn) error {
		if platform.ShouldBlockConnection() {
			return errBlocked
		}

		return conn.Control(func(fd uintptr) {
			app.MarkSocket(int(fd))
		})
	}
}
