package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger" // swagger handler
	"github.com/mhmmdFsl/go-fiber-rest-api-example/config"
	_ "github.com/mhmmdFsl/go-fiber-rest-api-example/docs"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/handler"
)

// @title Go Fiber Example API
// @version 1.0
// @description This is a sample for Fiber Rest API
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	app := fiber.New(config.FiberConfig())

	handlerCfg := &handler.HandlerCfg{
		App: app,
	}
	handler.NewHandler(handlerCfg)

	app.Listen(":3000")
}
