package middleware

import "github.com/gofiber/fiber/v2/middleware/basicauth"

var Basicauth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"john":  "doe",
		"admin": "123456",
	},
})
