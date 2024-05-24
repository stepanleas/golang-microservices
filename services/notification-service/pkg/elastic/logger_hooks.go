package elastic

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/go-extras/elogrus.v8"
)

type ElasticLoggerConfig struct {
	Client *ElasticClient
	Host   string
	Level  logrus.Level
	Index  string
}

func LoggerHook(cfg ElasticLoggerConfig) (*elogrus.ElasticHook, error) {
	hook, err := elogrus.NewAsyncElasticHook(cfg.Client.client, cfg.Host, cfg.Level, cfg.Index)
	if err != nil {
		return nil, err
	}

	return hook, nil
}
