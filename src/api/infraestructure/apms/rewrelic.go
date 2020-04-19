package apms

import (
	newrelic "github.com/newrelic/go-agent"
	"log"
	"os"
)

func NewNewRelicApp() newrelic.Application {
	newRelicApp, err := newrelic.NewApplication(
		newrelic.Config{AppName: "safe-taiga-63543",
			License: os.Getenv("NEW_RELIC_LICENSE_KEY"),
			Enabled:true})
	if err != nil {
		log.Println("Could not init rewrelic agent")
		// Returns nil for localhost tests
		return nil
	}
	return newRelicApp
}
