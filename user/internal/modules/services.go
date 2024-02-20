package modules

import (
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/infrastructure/component"
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/modules/user/service"
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/storages"
)

type Services struct {
	User service.Userer
}

func NewServices(storages *storages.Storages, components *component.Components) *Services {
	userService := service.NewUser(storages.User, components.Logger)

	return &Services{
		User: userService,
	}
}
