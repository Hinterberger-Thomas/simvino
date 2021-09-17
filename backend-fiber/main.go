package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/Hinterberger-Thomas/simvino/api"
	"github.com/Hinterberger-Thomas/simvino/db"
	"github.com/Hinterberger-Thomas/simvino/handler"
	"github.com/Hinterberger-Thomas/simvino/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	db.InitPostgres()
	db.InitMongo()
	go api.GetCoinsValue()

	auth := app.Group("/auth", logger.New(), middleware.Basicauth)
	userApi := app.Group("/user", logger.New(), middleware.Basicauth, middleware.ValSession)
	app.Get("/", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return nil
	})
	app.Get("/swagger/*", swagger.Handler)

	auth.Get("/login", handler.HandlerLogin)

	auth.Get("/register", handler.HandleRegister)

	userApi.Get("/addCurrency", handler.HandleAddCurrency)

	ln, _ := net.Listen("tcp", "localhost:3000")

	cer, _ := tls.LoadX509KeyPair("./host.cert", "./host.key")

	ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})

	log.Fatal(app.Listener(ln))
}
