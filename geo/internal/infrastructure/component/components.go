package component

import (
	"projects/LDmitryLD/hugoproxy-microservices/geo/config"

	"go.uber.org/zap"
)

type Components struct {
	Conf config.AppConf

	Logger *zap.Logger
}

func NewComponents(conf config.AppConf, logger *zap.Logger) *Components {
	return &Components{
		Conf:   conf,
		Logger: logger,
	}
}
