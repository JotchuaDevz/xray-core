package udp

import (
	"github.com/JotchuaDevz/xray-core/common"
	"github.com/JotchuaDevz/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
