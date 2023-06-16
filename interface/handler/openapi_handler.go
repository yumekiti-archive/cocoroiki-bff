package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type OpenAPIHandler interface {
	OpenAPIHandler(c echo.Context) error
}

type openapiHandler struct{}

func NewOpenAPIHandler() OpenAPIHandler {
	return &openapiHandler{}
}

func (h *openapiHandler) OpenAPIHandler(c echo.Context) error {
	OpenAPIURL := "https://cocoroiki-moc.yumekiti.net/api"

	req, err := http.Get(OpenAPIURL + strings.Replace(c.Request().URL.Path, "/mock", "", 1))
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(req.StatusCode, body)
}