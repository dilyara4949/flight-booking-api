package config

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]string
		want    Config
		wantErr error
	}{
		{
			name: "OK",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			want: Config{
				Address:           "address",
				RestPort:          "restport",
				JWTTokenSecret:    "secret",
				AccessTokenExpire: 24,
				PostgresCfg: PostgresCfg{
					Host:           "postgreshost",
					Port:           "postgresport",
					User:           "postgresuser",
					Password:       "postgrespass",
					Name:           "postgresdbname",
					Timeout:        1,
					MaxConn:        2,
					ContextTimeout: time.Second * time.Duration(23),
				},
			},
		},
		{
			name: "empty rest port",
			input: map[string]string{
				"ADDRESS":                  "address",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingRestPort,
		},
		{
			name: "empty address",
			input: map[string]string{
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingAddress,
		},
		{
			name: "empty jwt secret",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingJWTTokenSecret,
		},
		{
			name: "empty postgres host",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresHost,
		},
		{
			name: "empty postgres port",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresPort,
		},
		{
			name: "empty postgres user",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresUser,
		},
		{
			name: "empty postgres password",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresPassword,
		},
		{
			name: "empty postgres db name",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresName,
		},
		{
			name: "empty postgres timeout",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresTimeout,
		},
		{
			name: "empty postgres max connections",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errMissingPostgresMaxConn,
		},
		{
			name: "empty access token expire",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"POSTGRES_CONTEXT_TIMEOUT": "23",
			},
			wantErr: errAccessTokenExpire,
		},
		{
			name: "empty context timeout",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"ACCESS_TOKEN_EXPIRE":      "24",
				"POSTGRES_MAX_CONNECTIONS": "2",
			},
			wantErr: errMissingPostgresContextTimeout,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.input {
				t.Setenv(k, v)
			}

			got, err := NewConfig()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
