// Package config implements go-config with env and k8s configmap
package config

import (
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap/v2"
)

// NewConfig returns config with env and k8s configmap setup
func NewConfig(opts ...config.Option) config.Config {
	cfg, _ := config.NewConfig()
	cfg.Load(
		env.NewSource(),
		configmap.NewSource(),
	)
	return cfg
}
