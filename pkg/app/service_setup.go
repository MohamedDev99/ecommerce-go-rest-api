package app

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/controllers"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/middleware"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/routes"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db"
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/repos"
	"github.com/gofiber/fiber/v3"
)

var log = configs.Logger

// setup the app
func ServiceSetup(app *fiber.App) {
	// read config
	log.Info("📝 Reading config file ... ")
	configs.LoadEnvFromDotEnv()

	// connect to database
	log.Info("🗃️ Connecting to database ... ")
	mongodb_client := db.ConnectToMongoDB()

	// configure services
	log.Info("🙈 Configure services ... ")
	// auth service
	authRepository := repos.NewAuthRepoAdapter(mongodb_client)
	authController := controllers.NewAuthController(authRepository)
	authRouter := routes.NewAuthRouter(*authController)
	log.Info("---> ✅ Finish config auth service ...  ")

	// user service
	userRepository := repos.NewUserRepoAdapter(mongodb_client)
	userController := controllers.NewUserController(userRepository)
	userRouter := routes.NewUserRouter(*userController)
	log.Info("---> ✅ Finish config user service ...  ")

	// product service
	productRepository := repos.NewProductRepoAdapter(mongodb_client)
	productController := controllers.NewProductController(productRepository) // TODO : product controller
	productRouter := routes.NewProductRouter(*productController)
	log.Info("---> ✅ Finish config product service...  ")

	// cart service
	cartRepository := repos.NewCartRepoAdapter(mongodb_client)
	cartController := controllers.NewCartController(cartRepository) // TODO : cart controller
	cartRouter := routes.NewCartRouter(*cartController)
	log.Info("---> ✅ Finish config cart service...  ")

	// address service
	addressRepository := repos.NewAddressRepoAdapter(mongodb_client)
	addressController := controllers.NewAddressController(addressRepository) // TODO : address controller
	addressRouter := routes.NewAddressRouter(*addressController)
	log.Info("---> ✅ Finish config address service...  ")

	// finish services configuration
	log.Info("✅ Finish services configuration ... ")

	// middlewares
	middleware.MiddlewareSetup(app)

	// add routes
	log.Info("🚗 Adding routes ... ")
	apiVersionGroup := app.Group(configs.EnvAPIVersion()) // group by api version
	authRouter.Setup(apiVersionGroup)                     // add auth routes
	userRouter.Setup(apiVersionGroup)                     // add user routes
	productRouter.Setup(apiVersionGroup)                  // add product routes
	cartRouter.Setup(apiVersionGroup)                     // add cart routes
	addressRouter.Setup(apiVersionGroup)                  // add address routes

	// start server
	log.Info("🎬 Starting server ... ")
}
