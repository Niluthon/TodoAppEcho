package main

import (
	"TodoAppEcho/app"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()
	e.Validator = &app.CustomValidator{Validator: validator.New()}

	RegisterRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
