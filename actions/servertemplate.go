package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// ServertemplateIndex default implementation.
func ServertemplateIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("servertemplate/index.html"))
}

