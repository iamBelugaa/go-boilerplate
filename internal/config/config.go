// Package config provides the application's centralized configuration system.
//
// Prefix naming convention:
//   - All environment variables must be prefixed with your service/application name.
//   - For a real application, replace "BOILERPLATE_" with your actual service name
//     (e.g., "MYAPP_", "PAYMENTS_", "ORDERS_") to avoid collisions.
package config

import (
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"

	"github.com/iamBelugaa/go-boilerplate/pkg/validation"
)

// LoadFromEnv loads application configuration from environment variables.
func LoadFromEnv() (*Config, error) {
	k := koanf.New(".")

	if err := k.Load(
		env.Provider(
			"BOILERPLATE_", ".",
			func(s string) string {
				return strings.ToLower(strings.TrimPrefix(s, "BOILERPLATE_"))
			},
		),
		nil,
	); err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := k.Unmarshal("", conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// Validate checks the loaded configuration for correctness.
func Validate(conf *Config) error {
	if err := validation.Check(conf); err != nil {
		return err
	}

	if conf.Server != nil {
		if err := conf.Server.Validate(); err != nil {
			return err
		}
	}

	if conf.Logging != nil {
		if err := conf.Logging.Validate(); err != nil {
			return err
		}
	}

	if conf.Database != nil {
		if err := conf.Database.Validate(); err != nil {
			return err
		}
	}

	if conf.Service != nil {
		if err := conf.Service.Validate(); err != nil {
			return err
		}
	}

	if conf.HealthChecks != nil {
		if err := conf.HealthChecks.Validate(); err != nil {
			return err
		}
	}

	return nil
}
