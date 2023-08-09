package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	MySqlDatabase string `env:"MYSQL_DATABASE"`
	MySqlUsername string `env:"MYSQL_USERNAME"`
	MySqlPassword string `env:"MYSQL_PASSWORD"`
	MySqlPort     string `env:"MYSQL_PORT"`
	MySqlHost     string `env:"MYSQL_HOST"`

	ServiceURL string `env:"SERVICE_URL"`

	MinioHost       string `env:"MINIO_HOST"`
	MinoAccessKey   string `env:"MINIO_ACCESS_KEY"`
	MinoSecretKey   string `env:"MINIO_SECRET_KEY"`
	MinioUseSSL     bool   `env:"MINIO_USE_SSL"`
	MinioBucketName string `env:"MINIO_BUCKET_NAME"`

	UserServiceURL string `env:"USER_SERVICE_URL"`
	HttpPort       string `env:"HTTP_PORT"`
}

func LoadConfig() (config Config) {
	if err := env.Parse(&config); err != nil {
		panic(any(err))
	}
	return config
}
