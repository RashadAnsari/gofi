package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func Init(app string, fileName string, cfg interface{}, defaultConfig string, prefix string) error {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(defaultConfig))); err != nil {
		return err
	}

	v.SetConfigName(fileName)
	v.SetEnvPrefix(prefix)
	v.AddConfigPath(fmt.Sprintf("/etc/%s/", app))
	v.AddConfigPath(fmt.Sprintf("$HOME/.%s", app))
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	switch err := v.MergeInConfig(); err.(type) {
	case nil:
	case viper.ConfigFileNotFoundError:
		logrus.Warn("no config file found. using defaults and environment variables")
	default:
		return err
	}

	if err := v.UnmarshalExact(&cfg); err != nil {
		return err
	}

	return nil
}
