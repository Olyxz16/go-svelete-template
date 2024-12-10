package handlers

import (
	"github.com/labstack/echo/v4"
)

var (
    StaticFilepath string
)

func Index(c echo.Context) error {
    err := c.File(StaticFilepath + "/index.html")
    return err
}


