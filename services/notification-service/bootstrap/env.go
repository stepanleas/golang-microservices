package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                    string `mapstructure:"APP_ENV"`
	ServerAddress             string `mapstructure:"SERVER_ADDRESS"`
	RabbitMqUrl               string `mapstructure:"RABBITMQ_URL"`
	ElasticSearchUrl          string `mapstructure:"ELASTICSEARCH_URL"`
	ElasticSearchAmpServerUrl string `mapstructure:"ELASTICSEARCH_APM_SERVER_URL"`
	ElasticSearchIndex        string `mapstructure:"ELASTICSEARCH_INDEX"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
