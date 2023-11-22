package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	GRPC struct {
		Port int `env:"GRPC_PORT" env-default:"12201"`
	}
	REST struct {
		Port int `env:"HTTP_PORT" env-default:"8081"`
	}
	RusProfile struct {
		BaseUrl string `env:"BASE_URL" env-default:"https://www.rusprofile.ru"`
	}
	Swagger struct {
		FilePath string `env:"SWAGGER_FILE_PATH" env-default:"api/profile_v1/service.swagger.json"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatal(err)
		}

	})
	return instance
}
