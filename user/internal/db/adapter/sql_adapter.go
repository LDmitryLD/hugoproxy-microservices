package adapter

import (
	"database/sql"
	"projects/LDmitryLD/hugoproxy-microservices/user/internal/models"

	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=SQLAdapterer
type SQLAdapterer interface {
	GetByEmail(email string) (models.UserDTO, error)
	Insert(user models.UserDTO) error
	List() ([]models.User, error)
}

type SQLAdapter struct {
	db *sqlx.DB
}

func NewSQLAdapter(db *sqlx.DB) *SQLAdapter {
	return &SQLAdapter{
		db: db,
	}
}

func (s *SQLAdapter) GetByEmail(email string) (models.UserDTO, error) {
	q := `
	SELECT * FROM users WHERE email = $1	
	`
	var user models.UserDTO
	err := s.db.Get(&user, q, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserDTO{}, status.Error(codes.NotFound, "not found")
			//return models.UserDTO{}, errors.ErrNotFound
		}
		return models.UserDTO{}, status.Error(codes.Internal, err.Error())
	}

	return user, err
}

func (s *SQLAdapter) Insert(user models.UserDTO) error {
	q := `
	INSERT INTO users
		(name, email, password)
	VALUES
		($1, $2, $3)		
	`
	_, err := s.db.Exec(q, user.Name, user.Email, user.Password)

	return err
}

func (s *SQLAdapter) List() ([]models.User, error) {
	q := `
	SELECT 
		name, email
	FROM 
	 	users
	`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Name, &u.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
