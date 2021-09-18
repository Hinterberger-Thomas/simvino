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

// GetUsers ... Get all users
// @Param user body userPas true "Account ID"
// @Summary Insert User into the DB
// @Description Insert the user with argoin2id in the db. Adds salt and pepper in the process.
// @P
// @Tags test
// @Success 200 {object} userPas
// @Router /auth/addCurrency [get]
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

// @Summary hey
// @Description hello
// @Produce json
// @Param body body userPas true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/login [post]
func HandleRegister(c *fiber.Ctx) error {

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
