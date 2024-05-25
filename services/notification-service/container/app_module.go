package container

import (
	"github.com/stepanleas/notification-service/bootstrap"
	"github.com/stepanleas/notification-service/container/api"
	"go.uber.org/fx"
)

var ApplicationModule = fx.Module("app",
	fx.Provide(provideApplication),
	loggerModule,
	elasticSearchModule,
	rabbitMqModule,
	api.ApiModule,
)

func provideApplication() bootstrap.Application {
	return bootstrap.App()
}
