package middleware

import (
	"fmt"

	"github.com/Hinterberger-Thomas/simvino/auth/paseto"
	"github.com/Hinterberger-Thomas/simvino/model/user"
	"github.com/gofiber/fiber/v2"
)

var UserCtxKey = "user"

type contextKey struct {
	name string
}

func ValSession(c *fiber.Ctx) error {
	sess, err := Sessions.Get(c)
	if err != nil {
		return err
	}
	sessionPer := sess.Get("session")
	if sessionPer == nil {
		return fmt.Errorf("session does not exist on the server")
	}

	payload, err := paseto.PaseMaker.VerifyToken(fmt.Sprintf("%v", sessionPer))
	if err != nil {
		return fmt.Errorf("session does not exist on the server")
	}
	c.Locals(UserCtxKey, user.BasicUser{User_id: payload.UserId, Role: payload.Role})
	if err != nil {
		return err
	}

	err = sess.Regenerate()
	sess.Set("session", sessionPer)
	if err != nil {
		return err
	}
	err = sess.Save()
	if err != nil {
		return err
	}
	c.Next()
	return nil
}
