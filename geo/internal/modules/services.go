package modules

import (
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/infrastructure/component"
	geoservice "projects/LDmitryLD/hugoproxy-microservices/geo/internal/modules/geo/service"
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/storages"
)

type Services struct {
	Geo geoservice.Georer
}

func NewServices(storages *storages.Storages, components *component.Components) *Services {
	geoService := geoservice.NewGeo(storages.Geo, components.Logger)

	return &Services{
		Geo: geoService,
	}
}
