package main

import (
	"goapi/api"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", healthCheck)

	userGroup := e.Group("/api/users")
	api.HandleUserGroup(userGroup)

	log.Fatal(e.Start(":6969"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
