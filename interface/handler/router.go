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
	openapiHandler OpenAPIHandler,
) {
	e.POST("/signin", func(c echo.Context) error {
		return config.Login(c)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ok",
		})
	})

	e.GET("/mock/*", openapiHandler.OpenAPIHandler)

	// r := e.Group("")
	// r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	e.GET("/*", strapiHandler.StrapiHandler)
	e.POST("/*", strapiHandler.StrapiHandler)
	e.PUT("/*", strapiHandler.StrapiHandler)
	e.DELETE("/*", strapiHandler.StrapiHandler)
}