package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// GamerserverIndex default implementation.
func GamerserverIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("gamerserver/index.html"))
}

