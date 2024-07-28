package router

import (
	"github.com/labstack/echo/v4"
	"monsterXknight/app/controllers"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	routes := e.Group("/api/")

	routes.Any("monster", controllers.Get)

	return e
}
