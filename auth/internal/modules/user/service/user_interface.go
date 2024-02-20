package service

import (
	"context"
	"projects/LDmitryLD/hugoproxy-microservices/auth/internal/models"
)

type Userer interface {
	Profile(ctx context.Context, email string) (models.User, error)
	Create(ctx context.Context, user models.User) error
}

type ProfileIn struct {
	Email string
}
type PrfileOut struct {
	Name     string
	Email    string
	Password string
}

type CreateIn struct {
	Name     string
	Email    string
	Password string
}

type CreateOut struct {
	Success bool
}
