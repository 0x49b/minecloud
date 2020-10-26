package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// ServerIndex default implementation.
func ServerIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("server/index.html"))
}

