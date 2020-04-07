package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infraestructure/handlers"
)

type App struct {
	engine *gin.Engine
	statusHandler *handlers.StatusHandler
}

func NewApp() *App {
	app := &App{}
	app.engine = gin.Default()
	wire(app)
	mapRoutes(app)
	return app
}

func (a *App) Run(port string) error {
	return a.engine.Run(port)
}
