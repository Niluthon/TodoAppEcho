package tasks

import "github.com/labstack/echo/v5"

type Handler struct {
}

func (h *Handler) TestTaskHandler(c *echo.Context) error {
	return c.String(200, "Hello, World From Tasks Handler!")
}
