//+build wireinject

package wire

import (
	"github.com/smurdoch/go-rest-api/src/api/services"
	"github.com/smurdoch/go-rest-api/src/api/controllers"
	"github.com/google/wire"
)

func InitializeApp() app.App {
	wire.Build(
		app.New,
		controllers..NewStatusController,
		controllers.NewAdmissionController,
		services.NewAdmissionService,
		repositories.NewAdmissionRepository,
	)
	return app.App{}
}
