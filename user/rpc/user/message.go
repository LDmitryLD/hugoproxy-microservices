package user

import "projects/LDmitryLD/hugoproxy-microservices/user/internal/models"

type ProfileIn struct {
	Email string
}

type ProfileOut struct {
	Name     string
	Email    string
	Password string
}

type ListIn struct{}

type ListOut struct {
	Users []models.User
}

type CreateIn struct {
	Name     string
	Email    string
	Password string
}

type CreateOut struct {
	Success bool
}
