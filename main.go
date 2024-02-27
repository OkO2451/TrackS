package main

import (
	"log"
	"os"

	"github.com/OkO2451/TrackS/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func main() {

	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", handlers.HandleGetHome)

	log.Fatal(app.Listen(os.Getenv("PORT")))

}

func InitApp() (*fiber.App, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	engine := django.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		PassLocalsToViews: true,
		Views:             engine,
	})

	app.Static("/static", "./static")
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return app, nil
}
