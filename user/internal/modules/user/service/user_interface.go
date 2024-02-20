package service

import "projects/LDmitryLD/hugoproxy-microservices/user/internal/models"

type Userer interface {
	Profile(email string) (models.UserDTO, error)
	Create(user models.UserDTO) error
	List() ([]models.User, error)
}
