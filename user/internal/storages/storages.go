package storages

import (
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/db/adapter"
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/modules/user/storage"
)

type Storages struct {
	User storage.UserStorager
}

func NewStorages(sqlAdapter *adapter.SQLAdapter) *Storages {
	userStorage := storage.NewUserStorage(sqlAdapter)

	return &Storages{
		User: userStorage,
	}
}
