package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-svelte-template/database"
)

func Ping(c echo.Context) error {
    if database.New().Health() {
        return c.JSON(200, map[string]string {"message": "OK"})
    }
    return c.JSON(500, map[string]string {"message": "KO"})
}
