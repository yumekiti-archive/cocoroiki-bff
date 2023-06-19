package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StrapiHandler interface {
	StrapiHandler(c echo.Context) error
}

type strapiHandler struct{}

func NewStrapiHandler() StrapiHandler {
	return &strapiHandler{}
}

func (h *strapiHandler) StrapiHandler(c echo.Context) error {
	StrapiURL := "https://cocoroiki-strapi.yumekiti.net"

	q := c.Request().URL.Query()
	q.Add("populate", "*")
	c.Request().URL.RawQuery = q.Encode()

	req, err := http.Get(StrapiURL + c.Request().URL.Path + "?" + c.Request().URL.RawQuery)
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
