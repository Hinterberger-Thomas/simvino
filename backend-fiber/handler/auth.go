package handler

import (
	"fmt"

	"github.com/Hinterberger-Thomas/simvino/auth"
	"github.com/Hinterberger-Thomas/simvino/auth/paseto"
	"github.com/Hinterberger-Thomas/simvino/middleware"
	"github.com/Hinterberger-Thomas/simvino/model/user"
	"github.com/Hinterberger-Thomas/simvino/model/userSession"
	"github.com/gofiber/fiber/v2"
)

type userPas struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUser godoc
// @Summary Get user account detail
// @Router /user/detail [gets]
func HandlerLogin(c *fiber.Ctx) error {
	userPas := userPas{}

	if err := c.QueryParser(&userPas); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId, passwordDb, err := user.GetIdAndPassowrdByEmail(userPas.Email)
	if err != nil {
		return err
	}
	hashed, err := auth.GeneratePassword(auth.DefaultConfig, userPas.Password)
	if err != nil {
		return err
	}
	if err == nil && passwordDb == hashed && userId != 0 {
		store, err := middleware.Sessions.Get(c)
		if err != nil {
			return err
		}
		token, err := paseto.PaseMaker.CreateToken(paseto.PayloadGet{UserId: userId, Role: ""})
		if err != nil {
			return err
		}

		store.Set("session", token)
		store.Save()

	} else {
		return fmt.Errorf("wrong email or password")
	}
	c.SendStatus(200)
	return err
}

var HandleRegister = func(c *fiber.Ctx) error {

	userPas := userPas{}
	if err := c.QueryParser(&userPas); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId, err := user.InserUser(user.User{Email: userPas.Email, Password: userPas.Password, Role: ""})
	if err == nil {
		store, err := middleware.Sessions.Get(c)
		sessoinId := store.ID()
		if err != nil {
			return err
		}
		token, err := paseto.PaseMaker.CreateToken(paseto.PayloadGet{UserId: userId, Role: ""})
		if err != nil {
			return err
		}

		store.Set("session", token)
		err = store.Save()
		if err != nil {
			return err
		}
		userSession.InsertSession(userSession.Users_Sessions_Ins{Fk_user_id: userId, SESSION_ID: sessoinId})
		c.SendStatus(200)
		return nil
	}
	c.SendStatus(921)
	return nil
}
