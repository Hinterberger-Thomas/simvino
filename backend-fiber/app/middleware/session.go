package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var (
	Storage = redis.New(redis.Config{
		Host:     "localhost",
		Port:     6380,
		Username: "",
		Password: "",
		Database: 0,
		URL:      "",
		Reset:    false,
	})

	Sessions = session.New(session.Config{
		Storage:        Storage,
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
	})
)
