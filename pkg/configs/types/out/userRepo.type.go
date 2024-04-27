package types

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"
)

// user repository
type UserRepository interface {
	CreateUser(user models.User) error
	FindAllUsers() ([]models.User, error)
	// FindUserByEmail(email string) (models.User, error)
	FindUserByUserId(userId string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(userId string) error
}

// auth repository
type AuthRepository interface {
	Login(email string) (models.User, error)
	Register(user models.User) error
}
