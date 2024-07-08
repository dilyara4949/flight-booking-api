package config

type Config struct {
	JWTTokenSecret string
	RestPort       string
	Address        string
	PostgresCfg
}

type PostgresCfg struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Timeout  int
	MaxConn  int
}
