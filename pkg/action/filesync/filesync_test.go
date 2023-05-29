package filesync_test

import (
	"godeploy/pkg/action"
	"godeploy/pkg/action/filesync"
)

var _ action.IAction = (*filesync.FileSync)(nil)
