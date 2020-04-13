package http

import (
	newrelic "github.com/newrelic/go-agent"
	"log"
	"os"
)

func rewrelic(app *App) {
	newRelicApp, err := newrelic.NewApplication(
		newrelic.Config{AppName: "safe-taiga-63543",
			License: os.Getenv("NEW_RELIC_LICENSE_KEY"),
			Enabled:true})
	if err != nil {
		log.Fatal("Could not init rewrelic agent")
	}
	app.newRelicApp = newRelicApp
}
