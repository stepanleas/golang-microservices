package container

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/stepanleas/notification-service/bootstrap"
	"github.com/stepanleas/notification-service/pkg/elastic"
	"github.com/stepanleas/notification-service/pkg/logger"
	"go.uber.org/fx"
	"gopkg.in/go-extras/elogrus.v8"
)

var loggerModule = fx.Module("logger",
	fx.Provide(provideLogger),
	fx.Invoke(addElasticHook),
)

func provideLogger() *logger.LogrusLogger {
	logrusLogger, err := logger.NewLogrusLogger()
	if err != nil {
		log.Fatal("could not create logger!")
	}

	return logrusLogger
}

func addElasticHook(
	client *elastic.ElasticClient,
	app bootstrap.Application,
	logger *logger.LogrusLogger,
) {
	hook, err := elogrus.NewAsyncElasticHook(
		client.Client(),
		app.Env.ElasticSearchUrl,
		logrus.DebugLevel,
		app.Env.ElasticSearchIndex,
	)
	if err != nil {
		log.Fatal("could not create elastic logger hook!")
	}

	logger.AddHook(hook)
}
