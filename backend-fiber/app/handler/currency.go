package handler

import (
	"github.com/Hinterberger-Thomas/simvino/middleware"
	"github.com/Hinterberger-Thomas/simvino/model/user"
	"github.com/Hinterberger-Thomas/simvino/model/userAssets"
	"github.com/gofiber/fiber/v2"
)

type asset struct {
	Crypto_id uint32 `json:"crypto_id"`
	Amount    uint32 `json:"amount"`
}

// fsadf ... Gfdasadfsfdas
// @Summary asdfasfasdfsadf
// @Description asdfasfdasf
// @P
// @Tags tests
// @Router /user/addCurrency [post]
func HandleAddCurrency(c *fiber.Ctx) error {
	user := c.Locals(middleware.UserCtxKey).(user.BasicUser)

	assetParam := asset{}
	if err := c.QueryParser(&assetParam); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := userAssets.InserUserAsset(userAssets.User_assets_ins{Crypto_id: assetParam.Crypto_id, Amount: assetParam.Amount, Fk_user_id: user.User_id})
	if err != nil {
		return err
	}
	c.SendStatus(200)
	return nil
}
