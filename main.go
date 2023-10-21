package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger" // swagger handler
	"github.com/mhmmdFsl/go-fiber-rest-api-example/config"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/db"
	_ "github.com/mhmmdFsl/go-fiber-rest-api-example/docs"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/handler"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/repository"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/service"
)

// @title			Go Fiber Example API
// @version		1.0
// @description	This is a sample for Fiber Rest API
// @termsOfService	http://swagger.io/terms/
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {

	app := fiber.New(config.FiberConfig())

	db := db.NewDb()
	catRepo := repository.NewCatRepository(db)
	catService := service.NewCatService(&service.CatServiceCfg{
		CatRepo: catRepo,
	})
	handlerCfg := &handler.HandlerCfg{
		App:        app,
		CatService: catService,
	}

	handler.NewHandler(handlerCfg)

	app.Listen(":3000")
}
