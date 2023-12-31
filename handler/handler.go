package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/service"
)

type handler struct {
	App        *fiber.App
	CatService service.CatService
}

type HandlerCfg struct {
	App        *fiber.App
	CatService service.CatService
}

func NewHandler(cfg *HandlerCfg) {

	h := &handler{
		App:        cfg.App,
		CatService: cfg.CatService,
	}

	// handle swagger request
	sw := h.App.Group("/swagger")
	sw.Get("*", swagger.HandlerDefault)

	// handle cat request
	cat := h.App.Group("/api/v1/cat")
	cat.Get("/hello", h.helloCat)
	cat.Post("/add", h.addCat)
	cat.Get("/:id", h.getCatById)
	cat.Get("/", h.getAllCat)
	cat.Delete("/:id", h.deleteCat)
}
