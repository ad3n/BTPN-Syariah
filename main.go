package main

import (
	"fmt"

	"github.com/ad3n/resto/configs"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/routes"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/ad3n/resto/docs"
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

// @title BTPN Syariah Test Code
// @version 1.0
// @description This Repository is only used for testing purpose
// @termsOfService http://swagger.io/terms/
// @contact.name Muhamad Surya Iksanudin
// @contact.email surya.iksanudin@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	app.Get("/docs/*", swagger.HandlerDefault)

	(routes.Menu{}).RegisterRoutes(app)
	(routes.Customer{}).RegisterRoutes(app)
	(routes.Order{}).RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%d", configs.Env.AppPort))
}
