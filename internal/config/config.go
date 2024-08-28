package config

import (
	"fmt"
	"time"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	JWTTokenSecret    string
	AccessTokenExpire int
	RestPort          string
	Address           string
	HeaderTimeout     time.Duration
	Postgres
	Kafka
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

type Kafka struct {
	Host      string
	Port      string
	EmailPush string
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Init(&cfg); err != nil {
		return Config{}, fmt.Errorf("get configs: %w", err)
	}

	return cfg, nil
}
