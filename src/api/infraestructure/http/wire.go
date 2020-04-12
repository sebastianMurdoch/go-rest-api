package http

import (
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users/service"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infraestructure/handlers"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infraestructure/persistence"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infraestructure/repositories"
	"log"
)

func wire(app *App)  {
	log.Printf("%s", "Init Wiring")

	/* Status Service*/
	app.statusHandler = handlers.NewStatusHandler(app.newRelicApp)

	/* Users Service */
	db := persistence.NewDB()
	uRepository := repositories.NewUserRepositoryImpl(db)
	uService := service.NewUserServiceImpl(uRepository)
	app.usersHandler = handlers.NewUsersHandler(uService, app.newRelicApp)


	log.Printf("%s", "Finish Wiring")
}
