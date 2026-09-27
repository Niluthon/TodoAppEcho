package tasks

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Handler struct {
}

func (h *Handler) Get(c *echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World From Tasks Handler!")
}

func (h *Handler) FetchAll(c *echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World From Tasks Handler!")
}
