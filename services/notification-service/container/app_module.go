package container

import (
	"log"

	"github.com/stepanleas/notification-service/bootstrap"
	"github.com/stepanleas/notification-service/container/api"
	"github.com/stepanleas/notification-service/pkg/logger"
	"go.uber.org/fx"
)

var ApplicationModule = fx.Module("app-module",
	fx.Provide(provideApplication),
	fx.Provide(provideLogger),
	api.ApiModule,
	ElasticSearchModule,
)

func provideApplication() bootstrap.Application {
	return bootstrap.App()
}

func provideLogger() *logger.LogrusLogger {
	logrusLogger, err := logger.NewLogrusLogger()
	if err != nil {
		log.Fatal("could not create logger!")
	}

	return logrusLogger
}
