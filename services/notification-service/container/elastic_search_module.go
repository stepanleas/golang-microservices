package container

import (
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
	"github.com/stepanleas/notification-service/bootstrap"
	"github.com/stepanleas/notification-service/pkg/elastic"
	"github.com/stepanleas/notification-service/pkg/logger"
	"go.uber.org/fx"
)

var ElasticSearchModule = fx.Module("elasticsearch",
	fx.Provide(provideClient),
	fx.Invoke(addElasticHook),
)

func provideClient(app bootstrap.Application) *elastic.ElasticClient {
	cfg := elasticsearch.Config{
		Addresses: []string{app.Env.ElasticSearchUrl},
	}

	retryCfg := elastic.RetryConfig{
		MaxRetries: 5,
		Delay:      time.Second * 15,
	}

	elastiClient, err := elastic.NewElasticClient(cfg, retryCfg)
	if err != nil {
		log.Fatal("Cannot setup elastic client", err)
	}

	return elastiClient
}

func addElasticHook(client *elastic.ElasticClient, app bootstrap.Application, logger *logger.LogrusLogger) {
	cfg := elastic.ElasticLoggerConfig{
		Client: client,
		Host:   app.Env.ElasticSearchUrl,
		Level:  logrus.DebugLevel,
		Index:  app.Env.ElasticSearchIndex,
	}

	hook, err := elastic.LoggerHook(cfg)
	if err != nil {
		log.Fatal("could not create logger!", err)
	}

	logger.AddHook(hook)
}
