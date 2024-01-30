package main

import (
	mg "github.com/OrigamiWang/msd/manage/facade"
)

var (
	Manage mg.IFUser = &mg.UserManageFacade{}
)
