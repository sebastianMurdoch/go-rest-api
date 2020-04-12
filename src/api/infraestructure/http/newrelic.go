package http

import (
	newrelic "github.com/newrelic/go-agent"
	"log"
	"os"
)

func rewrelic(app *App) {
	newRelicApp, err := newrelic.NewApplication(newrelic.Config{AppName:"go-rest-api", License:os.Getenv("NEW_RELIC_LICENSE_KEY")})
	if err != nil {
		log.Fatal("Could not init rewrelic agent")
	}
	app.newRelicApp = newRelicApp
}
