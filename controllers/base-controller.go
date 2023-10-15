package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBaseUrlController(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome todo list app")
}
