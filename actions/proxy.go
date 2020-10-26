package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// ProxyIndex default implementation.
func ProxyIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("proxy/index.html"))
}

