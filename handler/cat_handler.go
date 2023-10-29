package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/httperror"
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
		var he *httperror.HttpError
		if errors.As(err, &he) {
			return ctx.Status(he.GetStatusCode()).JSON(he)
		}
		return ctx.Status(500).JSON(httperror.NewInternal(err))
	}

	return ctx.JSON(rs)
}

// getCatById godoc
// @Summary      Get cat by id
// @Description  Get cat by id
// @Tags         cat
// @Param id path int true "cat id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Cat
// @Router       /api/v1/cat/{id} [get]
func (h *handler) getCatById(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	rs, err := h.CatService.GetById(id)
	if err != nil {
		var he *httperror.HttpError
		if errors.As(err, &he) {
			return ctx.Status(he.GetStatusCode()).JSON(he)
		}
		return ctx.Status(500).JSON(httperror.NewInternal(err))
	}

	return ctx.JSON(rs)
}

// getAllCat godoc
// @Summary      Get all cat
// @Description  Get all cat
// @Tags         cat
// @Accept       json
// @Produce      json
// @Success      200  {object}  []entity.Cat
// @Router       /api/v1/cat [get]
func (h *handler) getAllCat(ctx *fiber.Ctx) error {
	cats, err := h.CatService.GetAll()
	if err != nil {
		return ctx.JSON(httperror.NewInternal(err))
	}

	return ctx.JSON(cats)
}
