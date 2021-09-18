package main

import (
	"log"

	"github.com/Hinterberger-Thomas/simvino/api"
	"github.com/Hinterberger-Thomas/simvino/config"
	"github.com/Hinterberger-Thomas/simvino/db"
	_ "github.com/Hinterberger-Thomas/simvino/docs"
	"github.com/Hinterberger-Thomas/simvino/handler"
	"github.com/Hinterberger-Thomas/simvino/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	db.InitPostgres()
	app := fiber.New()

	db.InitMongo()
	config.InitConf()
	go api.GetCoinsValue()

	auth := app.Group("/auth", logger.New(), middleware.Basicauth)
	userApi := app.Group("/user", logger.New(), middleware.Basicauth, middleware.ValSession)
	app.Get("/", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return nil
	})
	app.Get("/swagger/*", swagger.Handler)

	auth.Post("/login", handler.HandlerLogin)

	auth.Post("/register", handler.HandleRegister)

	// fsadf ... Gfdasadfsfdas
	// @Summary asdfasfasdfsadf
	// @Description asdfasfdasf
	// @P
	// @Tags Hallo
	// @Router /user/addCurrency [post]
	userApi.Post("/addCurrency", handler.HandleAddCurrency)

	// ln, _ := net.Listen("tcp", "localhost:3000")

	// cer, _ := tls.LoadX509KeyPair("./cert/host.cert", "./cert/host.key")

	// //ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})

	// log.Fatal(app.Listener(ln))
	log.Fatal(app.Listen(":8080"))
}
