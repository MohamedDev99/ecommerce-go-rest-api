package controllers

import (
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/out"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userRepo types.UserRepository
}

// new instance of user controller
func NewUserController(userRepo types.UserRepository) *UserController {
	// initialize user controller
	return &UserController{
		userRepo: userRepo,
	}
}

// login user
func (u *UserController) GetUsers(c fiber.Ctx) error {
	// get user
	// getUsers, err := u.userRepo.CreateUser(models.User{})
	return c.SendStatus(200)
}
