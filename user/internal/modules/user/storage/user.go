package storage

import (
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/db/adapter"
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=UserStorager
type UserStorager interface {
	GetByEmail(email string) (models.UserDTO, error)
	Create(user models.UserDTO) error
	List() ([]models.User, error)
}

type UserStorage struct {
	adapter adapter.SQLAdapterer
}

func NewUserStorage(sqlAdapter adapter.SQLAdapterer) UserStorager {
	return &UserStorage{
		adapter: sqlAdapter,
	}
}

func (u *UserStorage) GetByEmail(email string) (models.UserDTO, error) {
	return u.adapter.GetByEmail(email)
}

func (u *UserStorage) Create(user models.UserDTO) error {
	return u.adapter.Insert(user)
}

func (u *UserStorage) List() ([]models.User, error) {
	return u.adapter.List()
}
