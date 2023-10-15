package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PsqlDSN             string `yaml:"psql_dsn"`
	JwtSecretKey        string `yaml:"jwt_secret_key"`
	Logger              Logger `yaml:"logger"`
	OriginRemote        string `yaml:"origin_remote"`
	FtpConnectionString string `yaml:"ftp_connection_string"`
	FtpUserLogin        string `yaml:"ftp_user_login"`
	FtpUserPassword     string `yaml:"ftp_user_password"`
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
