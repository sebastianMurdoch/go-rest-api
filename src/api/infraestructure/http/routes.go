package http

func mapRoutes(app *App)  {
	app.Engine.GET("/ping", app.StatusHandler.Ping)
	app.Engine.GET("/users", app.UsersHandler.Get)
	app.Engine.POST("/users", app.UsersHandler.Post)
}
