package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// PlayerIndex default implementation.
func PlayerIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("player/index.html"))
}

