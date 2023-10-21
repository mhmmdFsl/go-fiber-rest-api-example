package config

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"time"
)

func FiberConfig() fiber.Config {

	readTimeOut, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIME_OUT"))
	appName := os.Getenv("APP_NAME")
	return fiber.Config{
		AppName:     appName,
		ReadTimeout: time.Duration(readTimeOut),
	}
}
