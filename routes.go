package main

import (
	"TodoAppEcho/features/tasks"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(echo *echo.Echo) {
	tasksHandler := &tasks.Handler{}

	tasksGroup := echo.Group("/tasks")
	tasksGroup.GET("/:id", tasksHandler.Get)
	tasksGroup.GET("", tasksHandler.FetchAll)
}

//func testHandler(c *echo.Context) error {
//	return c.String(200, "Hello, World!")
//}
