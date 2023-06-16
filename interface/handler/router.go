package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yumekiti/cocoroiki-bff/config"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo,
) {
	e.POST("/signin", func(c echo.Context) error {
		return config.Login(c)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ok",
		})
	})

	// 以下のルーティングはJWT認証が必要
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))


	// 以下は別サーバーへのプロキシ
	e.GET("/mock/*", func(c echo.Context) error {
		 fastAPIURL := "https://cocoroiki-moc.yumekiti.net/api"

		//  リクエスト送信
		req, err := http.Get(fastAPIURL + strings.Replace(c.Request().URL.Path, "/mock", "", 1))
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスボディを読み込む
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSONBlob(req.StatusCode, body)
	})
}