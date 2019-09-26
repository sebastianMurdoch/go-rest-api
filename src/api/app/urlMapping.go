package app

const NAMESPACE = "/moderations/official-brand-gate"

func mapUrls(app *App) {

	app.engine.GET("/ping", app.statusController.Ping)
	app.engine.GET("/admissions", app.admissionController.GetAll)

}
