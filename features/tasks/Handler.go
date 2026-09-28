package tasks

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Handler struct {
}

func (h *Handler) Get(c *echo.Context) error {
	taskDTO := new(TaskDTO)

	if err := c.Bind(taskDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(taskDTO); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, taskDTO)
}

func (h *Handler) FetchAll(c *echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World From Tasks Handler!")
}
