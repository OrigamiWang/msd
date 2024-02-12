package cli

import (
	auth "github.com/OrigamiWang/msd/auth/facade"
	conf "github.com/OrigamiWang/msd/conf-center/facade"
)

var (
	Auth auth.IFAuth = &auth.AuthFacade{}
	Conf conf.IFConf = &conf.ConfCenterFacade{}
)
