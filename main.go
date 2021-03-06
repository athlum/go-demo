package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.POST("/", func(c echo.Context) error {
		data, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			e.Logger.Error("error", err)
		}
		defer c.Request().Body.Close()
		e.Logger.Info("body", data)
		return c.String(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
