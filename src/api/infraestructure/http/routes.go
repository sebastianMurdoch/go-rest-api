package http

import "log"

func mapRoutes(app *App)  {
	log.Println("URL Mapping: initialized")

	app.engine.GET("/status", app.statusHandler.Get)

	log.Println("URL Mapping: done")
}
