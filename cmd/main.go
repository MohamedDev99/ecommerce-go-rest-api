package main

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {

	// @DESC : initializes the app
	// Start the server
	serverApp := fiber.New(fiber.Config{
		AppName:       "Ecommerce API",
		CaseSensitive: true,
		StrictRouting: true,
	})

	// @DESC : close MongoDB connection
	// defer db.CloseMongoDBConnection(db.DB_CLIENT)

	// TODO : app setup
	// services
	app.ServiceSetup(serverApp)

	// @DESC : app listen
	if err := serverApp.Listen(":5000"); err != nil {
		configs.Logger.Error("Error while starting server at port :5000 ", zap.String("error", err.Error()))
	}

}
