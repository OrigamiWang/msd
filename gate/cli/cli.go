package cli

import (
	conf "github.com/OrigamiWang/msd/conf-center/facade"
)

var (
	Conf conf.IFConf = &conf.ConfCenterFacade{}
)
