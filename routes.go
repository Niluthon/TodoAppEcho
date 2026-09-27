package main

import (
	"TodoAppEcho/features/tasks"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(echo *echo.Echo) {
	tasksHandler := &tasks.Handler{}

	tasks := echo.Group("/tasks")
	tasks.GET("", tasksHandler.TestTaskHandler)
}

//func testHandler(c *echo.Context) error {
//	return c.String(200, "Hello, World!")
//}
