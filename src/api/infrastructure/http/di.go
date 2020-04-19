package http

import (
	"github.com/gin-gonic/gin"
	di_injector "github.com/sebastianMurdoch/di-injector"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users/service"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infrastructure/apms"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infrastructure/handlers"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infrastructure/persistence"
	"github.com/sebastianMurdoch/go-rest-api/src/api/infrastructure/repositories"
	"log"
)

func buildCustomDiContainer() di_injector.DiContainer {
	/* Creates empty container */
	c := di_injector.NewDiContainer()

	/* Here you need to add all the infrastructure components that your project is going to need */
	err := c.AddToDependencies(
		persistence.NewDB(), // Adds the postgres DB
		&repositories.UserRepositoryImpl{},
		&service.UserServiceImpl{},
		&handlers.UsersHandler{},
		&handlers.StatusHandler{},
		apms.NewNewRelicApp(),
		gin.Default(),
		)
	if err != nil {
		log.Fatal("Could not add dependencies to the container ")
	}

	return c
}
