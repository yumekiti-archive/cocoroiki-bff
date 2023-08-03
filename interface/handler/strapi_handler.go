package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"time"

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

	if c.Request().URL.Path == "/api/posts" {
		// Parse the response data
		var responseData map[string]interface{}
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			log.Fatal(err)
		}

		// Update createdAt and updatedAt fields
		data := responseData["data"].([]interface{})
		for _, v := range data {
			attributes := v.(map[string]interface{})["attributes"].(map[string]interface{})
			// add 9 hours to createdAt and updatedAt
			createdAt := attributes["createdAt"].(string)
			t, err := time.Parse(time.RFC3339, createdAt)
			if err != nil {
				log.Fatal(err)
			}
			attributes["createdAt"] = t.Add(9 * time.Hour).Format(time.RFC3339)

			updatedAt := attributes["updatedAt"].(string)
			t, err = time.Parse(time.RFC3339, updatedAt)
			if err != nil {
				log.Fatal(err)
			}
			attributes["updatedAt"] = t.Add(9 * time.Hour).Format(time.RFC3339)
		}

		// Convert the updated data back to JSON
		updatedBody, err := json.Marshal(responseData)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSONBlob(req.StatusCode, updatedBody)
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