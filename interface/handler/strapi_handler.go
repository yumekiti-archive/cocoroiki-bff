package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StrapiHandler interface {
	GetHandler(c echo.Context) error
	PostHandler(c echo.Context) error
	PutHandler(c echo.Context) error
	DeleteHandler(c echo.Context) error
}

type strapiHandler struct{}

func NewStrapiHandler() StrapiHandler {
	return &strapiHandler{}
}

func (h *strapiHandler) GetHandler(c echo.Context) error {
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

func (h *strapiHandler) PostHandler(c echo.Context) error {
	StrapiURL := "https://cocoroiki-strapi.yumekiti.net"

	req, err := http.Post(StrapiURL+c.Request().URL.Path, "application/json", c.Request().Body)
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

func (h *strapiHandler) PutHandler(c echo.Context) error {
	StrapiURL := "https://cocoroiki-strapi.yumekiti.net"

	req, err := http.NewRequest(http.MethodPut, StrapiURL+c.Request().URL.Path, c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(res.StatusCode, body)
}


func (h *strapiHandler) DeleteHandler(c echo.Context) error {
	StrapiURL := "https://cocoroiki-strapi.yumekiti.net"

	req, err := http.NewRequest(http.MethodDelete, StrapiURL+c.Request().URL.Path, c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSONBlob(res.StatusCode, body)
}