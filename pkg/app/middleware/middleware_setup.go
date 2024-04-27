package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

// setup middleware
func MiddlewareSetup(app *fiber.App) {

	// cors
	// app.Use(cors.New())
	// Logging Request ID
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "${time} --- from:[${ip}]:${port} - ${pid} ${locals:requestid} | ${status} | ${latency} | ${method} | ${path}​\u200b\n",
		TimeFormat: "2006-Jan-02 15:04:05 -07:00",
	}))

}
