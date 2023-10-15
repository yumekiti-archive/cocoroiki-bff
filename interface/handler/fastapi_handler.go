package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FastapiHandler interface {
	GetAppUsersHandler(c echo.Context) error
}

type fastapiHandler struct{}

func NewFastapiHandler() FastapiHandler {
	return &fastapiHandler{}
}

const FastapiURL = "http://cocoroiki-api:8080"

func (h *fastapiHandler) GetAppUsersHandler(c echo.Context) error {
	req, err := http.Get(FastapiURL + "/api/app-users")
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
