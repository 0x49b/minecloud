package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/lichtwellenreiter/minecloud/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
