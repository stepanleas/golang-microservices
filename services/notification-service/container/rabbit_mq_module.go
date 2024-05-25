package container

import (
	"log"

	"github.com/stepanleas/notification-service/bootstrap"
	"github.com/stepanleas/notification-service/pkg/rabbit_mq"
	"go.uber.org/fx"
)

var rabbitMqModule = fx.Module("rabbit-mq",
	fx.Provide(provideRabbitMqClient),
)

func provideRabbitMqClient(app bootstrap.Application) *rabbit_mq.RabbitMqClient {
	conn, err := rabbit_mq.NewRabbitMqClient(app.Env.RabbitMqUrl)

	if err != nil {
		log.Fatal("module: cannot connect to rabbitmq!")
	}

	return conn
}
