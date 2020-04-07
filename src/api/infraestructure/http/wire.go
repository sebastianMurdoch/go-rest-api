package http

import (
	statusService "github.com/sebastianMurdoch/go-rest-api/src/api/domain/status/service"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infraestructure/handlers"
	"log"
)

func wire(app *App)  {
	log.Printf("%s", "Init Wiring")

	/* Status Service*/
	statusService := statusService.NewItemsServiceImpl()
	app.statusHandler = handlers.NewStatusHandler(statusService)

	log.Printf("%s", "Finish Wiring")
}
