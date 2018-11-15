package grifts

import (
	"github.com/PeterSkopal/web_dashboard/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
