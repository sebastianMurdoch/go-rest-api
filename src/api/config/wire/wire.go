//+build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/smurdoch/my_simple_rest_api/src/api/app"
	"github.com/smurdoch/my_simple_rest_api/src/api/controllers"
	"github.com/smurdoch/my_simple_rest_api/src/api/repositories"
	"github.com/smurdoch/my_simple_rest_api/src/api/services"
)

func InitializeApp() app.App {
	wire.Build(
		app.New,
		controllers.NewStatusController,
		controllers.NewAdmissionController,
		services.NewAdmissionService,
		repositories.NewAdmissionRepository,
	)
	return app.App{}
}
