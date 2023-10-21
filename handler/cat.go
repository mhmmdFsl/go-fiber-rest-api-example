package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model"
)

// helloCat godoc
// @Summary      Say hello from cat
// @Description  Get greeting from cat
// @Tags         cat
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /api/v1/cat/hello [get]
func (h *handler) helloCat(c *fiber.Ctx) error {
	return c.SendString("Hello from cat!")
}

// addCat godoc
// @Summary      Add new cat
// @Description  Add new cat data
// @Tags         cat
// @Param Cat body model.AddCatDto true "Request Body"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Cat
// @Router       /api/v1/cat/add [post]
func (h *handler) addCat(ctx *fiber.Ctx) error {

	cat := new(model.AddCatDto)
	if err := ctx.BodyParser(cat); err != nil {
		return err
	}

	rs, err := h.CatService.AddNewCat(cat)
	if err != nil {
		if err != nil {
			return err
		}
	}

	return ctx.JSON(rs)
}
