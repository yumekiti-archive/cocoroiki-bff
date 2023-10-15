package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"

	"github.com/yumekiti/cocoroiki-bff/config"
)

func InitRouting(
	e *echo.Echo,
	strapiHandler StrapiHandler,
	fastapiHandler FastapiHandler,
) {
	e.POST("/signin", func(c echo.Context) error {
		return config.Login(c)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ok",
		})
	})

	// r := e.Group("")
	// r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	e.GET("/api/app-users", fastapiHandler.GetAppUsersHandler)
	e.GET("/*", strapiHandler.GetHandler)
	e.POST("/*", strapiHandler.PostHandler)
	e.PUT("/*", strapiHandler.PutHandler)
	e.DELETE("/*", strapiHandler.DeleteHandler)
}