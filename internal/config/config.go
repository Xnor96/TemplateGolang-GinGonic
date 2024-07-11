package config

import "os"

type Config struct {
	UriRedis      string
	UriMongo      string
	UriPostgresql string
}

func NewConfig() *Config {
	conf := &Config{}

	conf.UriRedis = os.Getenv("URI_REDIS")
	conf.UriMongo = os.Getenv("URI_MONGO")
	conf.UriPostgresql = os.Getenv("URI_POSTGRESQL")

	return conf

}
