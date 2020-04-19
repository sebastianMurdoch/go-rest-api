package http

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infrastructure/handlers"
	"log"
)

type App struct {
	NewRelicApp   newrelic.Application    `inject:"auto"`
	Engine        *gin.Engine             `inject:"auto"`
	StatusHandler *handlers.StatusHandler `inject:"auto"`
	UsersHandler  *handlers.UsersHandler  `inject:"auto"`
}

func NewApp() *App {
	app := &App{}

	/* Build custom container */
	c := buildCustomDiContainer()

	/* Wiring */
	err := c.InjectWithDependencies(app)
	if err != nil {
		log.Fatal("Could not inject dependencies")
	}

	/* Routing */
	mapRoutes(app)
	return app
}

func (a *App) Run(port string) error {
	return a.Engine.Run(port)
}
