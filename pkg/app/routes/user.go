package routes

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/controllers"
	"github.com/gofiber/fiber/v3"
)

type UserRouter struct {
	userController controllers.UserController
}

// new instance of user routes
func NewUserRouter(userController controllers.UserController) *UserRouter {
	// initialize user routes
	return &UserRouter{
		userController: userController,
	}
}

// user routes
func (u *UserRouter) Setup(userRouterGroup fiber.Router) {
	// routes
	userRoutes := userRouterGroup.Group("/users")
	// 1- user login
	userRoutes.Post("/login", u.userController.GetUsers)
	log.Info("-----> ✅ Finish user routes ... ")
}
