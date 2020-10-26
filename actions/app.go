package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
	"github.com/lichtwellenreiter/minecloud/models"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_minecloud_session",
		})

		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(csrf.New)
		app.Use(popmw.Transaction(models.DB))
		app.Use(translations())

		app.GET("/", func (c buffalo.Context) error {
			return c.Redirect(308, "dashboardPath()")
		}) //HomeHandler

		app.GET("/dashboard/index", DashboardIndex).Name("dashboard")
		app.GET("/dashboard/widget", DashboardWidget)
		app.GET("/player/index", PlayerIndex)
		app.GET("/proxy/index", ProxyIndex)
		app.GET("/server/index", ServerIndex)
		app.GET("/gamerserver/index", GamerserverIndex)
		app.GET("/servertemplate/index", ServertemplateIndex)
		app.GET("/serverplugin/index", ServerpluginIndex)
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
