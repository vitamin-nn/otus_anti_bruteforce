package config

import (
	"github.com/spf13/viper"
)

func Load(filepath string) (c *Config, err error) {
	viper.SetConfigFile(filepath)
	viper.SetConfigType("toml")
	setDefaults()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err = viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func setDefaults() {
	viper.SetDefault("log.log_level", "info")
}

type Config struct {
	Log        Log
	GrpcServer GrpcServer `mapstructure:"grpc_server"`
	PSQL       PSQL
	RateLimit  RateLimit
}

type PSQL struct {
	DSN string
}

type Log struct {
	LogLevel string `mapstructure:"log_level"`
}

type GrpcServer struct {
	Addr string
}

type RateLimit struct {
	Login    int
	Password int
	IP       int
	Duration int64
}