package app

import (
	"github.com/gin-gonic/gin"
	"github.com/smurdoch/my_simple_rest_api/src/api/controllers"
)

type App struct {
	engine              *gin.Engine
	statusController    *controllers.StatusController
	admissionController *controllers.AdmissionController
}

func New(statusController *controllers.StatusController, admissionController *controllers.AdmissionController) App {
	app := App{}
	app.statusController = statusController
	app.admissionController = admissionController
	return app
}

func (app *App) Start() {
	app.engine = gin.Default()
	mapUrls(app)
	app.engine.Run(":9090")
}
