package main

import (
	"fmt"

	"github.com/ad3n/resto/configs"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	configs.Load()
	configs.Db.AutoMigrate(
		&models.Customer{},
		&models.Menu{},
		&models.Order{},
		&models.OrderDetail{},
	)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	(routes.Menu{}).RegisterRoutes(app)
	(routes.Customer{}).RegisterRoutes(app)
	(routes.Order{}).RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%d", configs.Env.AppPort))
}
