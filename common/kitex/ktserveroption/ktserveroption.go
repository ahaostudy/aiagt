package ktserveroption

import (
	ktconf "github.com/aiagt/kitextool/conf"
	ktserver "github.com/aiagt/kitextool/suite/server"
	ktutils "github.com/aiagt/kitextool/utils"
)

type LocalIpOption struct {
	ktserver.EmptyOption
}

func (c LocalIpOption) Apply(_ *ktserver.KitexToolSuite, conf *ktconf.ServerConf) {
	conf.Server.Address = ktutils.CompleteAddress(conf.Server.Address)
}

func WithLocalIpOption() ktserver.Option {
	return LocalIpOption{}
}
