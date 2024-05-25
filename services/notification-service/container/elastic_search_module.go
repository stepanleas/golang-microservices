package container

import (
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/stepanleas/notification-service/bootstrap"
	"github.com/stepanleas/notification-service/pkg/elastic"
	"go.uber.org/fx"
)

var elasticSearchModule = fx.Module("elasticsearch",
	fx.Provide(provideElasticClient),
)

func provideElasticClient(app bootstrap.Application) *elastic.ElasticClient {
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
