package main

import (
	"log"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	RegisterRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
