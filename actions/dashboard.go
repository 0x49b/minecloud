package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// DashboardIndex default implementation.
func DashboardIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("dashboard/index.html"))
}

// DashboardWidget default implementation.
func DashboardWidget(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("dashboard/widget.html"))
}

