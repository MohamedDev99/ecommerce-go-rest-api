package routes

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/controllers"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs"
	"github.com/gofiber/fiber/v3"
)

// logger
var log = configs.Logger

type AuthRouter struct {
	AuthController controllers.AuthController
}

// create new auth router
func NewAuthRouter(authController controllers.AuthController) *AuthRouter {
	// initialize auth router
	return &AuthRouter{
		AuthController: authController,
	}
}

// setup auth routes
func (a *AuthRouter) Setup(authRouterGroup fiber.Router) {
	// routes
	authRoutes := authRouterGroup.Group("/auth")
	// 1- user login
	authRoutes.Post("/login", a.AuthController.Login)
	// 2- user register
	authRoutes.Post("/register", a.AuthController.Register)
	log.Info("-----> ✅ Finish auth routes ... ")
}
