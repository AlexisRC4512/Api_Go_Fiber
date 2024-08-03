package app

import (
	"fmt"

	_ "github.com/AlexisRC4512/Api_Go_Fiber/src/app/docs"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/handlers"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/middlewares"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title	Api go Fiber
// @version		1.0
// @description	This is an Api go Fiber for coding c
// @termsOfService	http://swagger.io/terms/
func Run() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	config.InitConfig()
	jwt := middlewares.NewAuthMiddleware(config.GetSecret())
	app.Post("/login", handlers.Login)
	app.Post("/getRotateMatriz", jwt, handlers.GetRotateMatrix)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Post("/factorize", jwt, handlers.FactorizeMatrix)
	port := config.GetPort()
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Println("Error al iniciar el servidor: %s", err)
	}
}
