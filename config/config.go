package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

type Config struct {
	Application Application
	DateBase    DateBase
}

type Application struct {
	Name    string
	Port string
}

type DateBase struct {
	Port string
}

func (c *Config) validate() error {
	return multierr.Combine(
		c.Application.validate(),
	)
}
func (a *Application) validate() error {
	if a.Port == "" {
		return errors.New("Application Port is empty")
	}
	return nil
}
func (d *DateBase) validate() error {
	if d.Port == "" {
		return errors.New("Db port is empty")
	}
	return nil
}
func Parse(filepath string) (*Config, error) {
	setDefaults()

	// Parse the file
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}

	// Unmarshal the config
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	// Validate the provided configuration
	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate the config")
	}
	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("Application.Port", "3010")
	viper.SetDefault("Application.Name", "grpc")

	viper.SetDefault("DateBase.Port", "6370")
}