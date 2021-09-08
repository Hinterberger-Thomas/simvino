package main

import (
	"log"
	"time"

	"github.com/Hinterberger-Thomas/simvino/auth"
	"github.com/Hinterberger-Thomas/simvino/db"
	"github.com/Hinterberger-Thomas/simvino/model/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()
	db.InitRedis()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	app.Post("/register", func(c *fiber.Ctx) error {
		payload := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		maker, err := auth.NewPasetoMaker("asdf")
		if err != nil {
			return err
		}
		token, err := maker.CreateToken(payload.Email, time.Duration(time.Now().Add(time.Hour*1000).Unix()))
		if err != nil {
			return err
		}
		user.InsertUserToken(token, payload.Email)
		return c.SendString("Hello, World!")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
