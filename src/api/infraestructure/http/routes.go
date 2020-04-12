package http

import (
	"log"
)

func mapRoutes(app *App)  {
	log.Println("URL Mapping: initialized")

	app.engine.GET("/ping", app.statusHandler.Ping)
	app.engine.GET("/users", app.usersHandler.Get)

	log.Println("URL Mapping: done")
}
