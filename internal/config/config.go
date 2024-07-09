package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

var (
	errMissingRestPort               = errors.New("REST_PORT is empty")
	errMissingAddress                = errors.New("ADDRESS is empty")
	errMissingJWTTokenSecret         = errors.New("JWT_TOKEN_SECRET is empty")
	errMissingPostgresHost           = errors.New("POSTGRES_HOST is empty")
	errMissingPostgresPort           = errors.New("POSTGRES_PORT is empty")
	errMissingPostgresName           = errors.New("POSTGRES_NAME is empty")
	errMissingPostgresUser           = errors.New("POSTGRES_USER is empty")
	errMissingPostgresPassword       = errors.New("POSTGRES_PASSWORD is empty")
	errMissingPostgresTimeout        = errors.New("POSTGRES_TIMEOUT is empty")
	errMissingPostgresMaxConn        = errors.New("POSTGRES_MAX_CONNECTIONS is empty")
	errPostgresMaxConnType           = errors.New("POSTGRES_MAX_CONNECTIONS must be an integer")
	errPostgresTimeoutType           = errors.New("POSTGRES_TIMEOUT must be an integer")
	errMissingPostgresRequestTimeout = errors.New("POSTGRES_REQUEST_TIMEOUT is empty")
	errPostgresRequestTimeoutType    = errors.New("POSTGRES_REQUEST_TIMEOUT must be an integer")
	errAccessTokenExpire             = errors.New("ACCESS_TOKEN_EXPIRE is empty")
	errAccessTokenExpireType         = errors.New("ACCESS_TOKEN_EXPIRE must be an integer")
)

const (
	jwtTokenSecretEnv         = "JWT_TOKEN_SECRET"
	restPortEnv               = "REST_PORT"
	addressEnv                = "ADDRESS"
	accessTokenExpireEnv      = "ACCESS_TOKEN_EXPIRE"
	postgresHostEnv           = "POSTGRES_HOST"
	postgresPortEnv           = "POSTGRES_PORT"
	postgresUserEnv           = "POSTGRES_USER"
	postgresPasswordEnv       = "POSTGRES_PASSWORD"
	postgresNameEnv           = "POSTGRES_DB"
	postgresTimeoutEnv        = "POSTGRES_TIMEOUT"
	postgresRequestTimeoutEnv = "POSTGRES_REQUEST_TIMEOUT"
	postgresMaxConnEnv        = "POSTGRES_MAX_CONNECTIONS"
)

type Config struct {
	JWTTokenSecret    string
	AccessTokenExpire int
	RestPort          string
	Address           string
	PostgresCfg
}

type PostgresCfg struct {
	Host           string
	Port           string
	User           string
	Password       string
	Name           string
	RequestTimeout time.Duration
	Timeout        int
	MaxConn        int
}

func NewConfig() (Config, error) {
	errs := make([]error, 0)

	jwtTokenSecret := os.Getenv(jwtTokenSecretEnv)
	if jwtTokenSecret == "" {
		errs = append(errs, errMissingJWTTokenSecret)
	}

	accessTokenExpireStr := os.Getenv(accessTokenExpireEnv)
	if accessTokenExpireStr == "" {
		errs = append(errs, errAccessTokenExpire)
	}

	accessTokenExpire, err := strconv.Atoi(accessTokenExpireStr)
	if err != nil {
		errs = append(errs, errAccessTokenExpire)
	}

	restPort := os.Getenv(restPortEnv)
	if restPort == "" {
		errs = append(errs, errMissingRestPort)
	}

	address := os.Getenv(addressEnv)
	if address == "" {
		errs = append(errs, errMissingAddress)
	}

	postgresHost := os.Getenv(postgresHostEnv)
	if postgresHost == "" {
		errs = append(errs, errMissingPostgresHost)
	}

	postgresPort := os.Getenv(postgresPortEnv)
	if postgresPort == "" {
		errs = append(errs, errMissingPostgresPort)
	}

	postgresUser := os.Getenv(postgresUserEnv)
	if postgresUser == "" {
		errs = append(errs, errMissingPostgresUser)
	}

	postgresPassword := os.Getenv(postgresPasswordEnv)
	if postgresPassword == "" {
		errs = append(errs, errMissingPostgresPassword)
	}

	postgresName := os.Getenv(postgresNameEnv)
	if postgresName == "" {
		errs = append(errs, errMissingPostgresName)
	}

	postgresTimeoutStr := os.Getenv(postgresTimeoutEnv)
	if postgresTimeoutStr == "" {
		errs = append(errs, errMissingPostgresTimeout)
	}

	postgresTimeout, err := strconv.Atoi(postgresTimeoutStr)
	if err != nil {
		errs = append(errs, errPostgresTimeoutType)
	}

	postgresRequestTimeoutStr := os.Getenv(postgresRequestTimeoutEnv)
	if postgresTimeoutStr == "" {
		errs = append(errs, errMissingPostgresRequestTimeout)
	}

	postgresRequstTimeout, err := strconv.Atoi(postgresRequestTimeoutStr)
	if err != nil {
		errs = append(errs, errPostgresRequestTimeoutType)
	}

	postgresMaxConnStr := os.Getenv(postgresMaxConnEnv)
	if postgresMaxConnStr == "" {
		errs = append(errs, errMissingPostgresMaxConn)
	}

	postgresMaxconn, err := strconv.Atoi(postgresMaxConnStr)
	if err != nil {
		errs = append(errs, errPostgresMaxConnType)
	}
	return Config{
		JWTTokenSecret:    jwtTokenSecret,
		RestPort:          restPort,
		Address:           address,
		AccessTokenExpire: accessTokenExpire,
		PostgresCfg: PostgresCfg{
			postgresHost,
			postgresPort,
			postgresUser,
			postgresPassword,
			postgresName,
			time.Hour * time.Duration(postgresRequstTimeout),
			postgresTimeout,
			postgresMaxconn,
		},
	}, nil
}
