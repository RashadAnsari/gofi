package config

import (
	"time"

	"github.com/RashadAnsari/gofi/pkg/s3"

	"github.com/RashadAnsari/gofi/pkg/config"
	"github.com/RashadAnsari/gofi/pkg/log"
)

const (
	app       = "gofi"
	cfgFile   = "config.yaml"
	cfgPrefix = "gofi"
)

type (
	Config struct {
		Logger log.Config `mapstructure:"logger"`
		Server Server     `mapstructure:"server"`
		S3     s3.Options `mapstructure:"s3"`
	}

	Server struct {
		Address         string        `mapstructure:"address"`
		ReadTimeout     time.Duration `mapstructure:"read-timeout"`
		WriteTimeout    time.Duration `mapstructure:"write-timeout"`
		GracefulTimeout time.Duration `mapstructure:"graceful-timeout"`
	}
)

func Init() (*Config, error) {
	var cfg Config

	if err := config.Init(app, cfgFile, &cfg, defaultConfig, cfgPrefix); err != nil {
		return nil, err
	}

	return &cfg, nil
}
