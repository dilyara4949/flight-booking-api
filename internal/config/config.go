package config

import (
	"fmt"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	JWTTokenSecret    string
	AccessTokenExpire int
	RestPort          string
	Address           string
	Postgres
}

type Postgres struct {
	Host           string
	Port           string
	User           string
	Password       string
	DB             string
	Timeout        int
	MaxConnections int
}

func NewConfig() (Config, error) {
	cfg := Config{}
	if err := envconfig.Init(&cfg); err != nil {
		return Config{}, fmt.Errorf("err=%s\n", err)
	}

	return cfg, nil
}
