package http

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infraestructure/handlers"
)

type App struct {
	newRelicApp newrelic.Application
	engine *gin.Engine
	statusHandler *handlers.StatusHandler
	usersHandler *handlers.UsersHandler
}

func NewApp() *App {
	app := &App{}
	app.engine = gin.Default()
	rewrelic(app)
	wire(app)
	mapRoutes(app)
	return app
}

func (a *App) Run(port string) error {
	return a.engine.Run(port)
}
