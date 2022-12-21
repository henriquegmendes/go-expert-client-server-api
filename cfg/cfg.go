package cfg

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Environment struct {
	Port                          int           `envconfig:"PORT"`
	DbQueryResponseTimeout        time.Duration `envconfig:"DB_QUERY_RESPONSE_TIMEOUT"`
	ExchangeApiBaseURL            string        `envconfig:"EXCHANGE_API_BASE_URL"`
	ExchangeApiResponseTimeout    time.Duration `envconfig:"EXCHANGE_API_RESPONSE_TIMEOUT"`
	InternalServerBaseURL         string        `envconfig:"INTERNAL_SERVER_BASE_URL"`
	InternalServerResponseTimeout time.Duration `envconfig:"INTERNAL_SERVER_RESPONSE_TIMEOUT"`
}

var env *Environment

func Env() *Environment {
	if env == nil {
		err := godotenv.Load()
		if err != nil {
			log.Printf("error loading .env file: %s", err.Error())
		}

		var newEnv Environment
		err = envconfig.Process("", &newEnv)
		if err != nil {
			log.Fatalf("error mapping envs into Environment struct: %s", err.Error())
		}

		env = &newEnv
	}

	return env
}
