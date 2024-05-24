package elastic

import (
	"time"

	"emperror.dev/errors"
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticClient struct {
	client *elasticsearch.Client
}

type RetryConfig struct {
	MaxRetries int
	Delay      time.Duration
}

// NewElasticClient creates a new ElasticClient with retry logic
func NewElasticClient(cfg elasticsearch.Config, retryCfg RetryConfig) (*ElasticClient, error) {
	var es *elasticsearch.Client
	var err error

	for i := 0; i < retryCfg.MaxRetries; i++ {
		es, err = elasticsearch.NewClient(cfg)
		if err == nil {
			client := &ElasticClient{client: es}
			if err = client.checkConnection(); err == nil {
				return client, nil
			}
		}

		time.Sleep(retryCfg.Delay)
	}

	return nil, errors.WrapIf(err, "v8.elasticsearch")
}

func (c *ElasticClient) checkConnection() error {
	timeout := time.Second * 10
	response, err := c.client.Cluster.Health(
		c.client.Cluster.Health.WithTimeout(timeout),
		c.client.Cluster.Health.WithPretty(),
	)

	if err != nil {
		return err
	}

	if response.IsError() {
		return errors.New("error checking elasticsearch connection")
	}

	return nil
}
