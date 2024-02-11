package cli

import (
	auth "github.com/OrigamiWang/msd/auth/facade"
)

var (
	Auth auth.IFAuth = &auth.AuthFacade{}
)
