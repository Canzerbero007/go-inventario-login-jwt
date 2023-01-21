package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/users")
	products := e.Group("/products")

	users.POST("/register", a.UserRegister)
	users.POST("/login", a.UserLogin)

	products.POST("/products", a.ProductRegister)
}
