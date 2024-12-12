package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/Olyxz16/go-svelte-template/database"
)


/*********************/
/* Healthcheck utils */
/*********************/

func Health(c echo.Context) error {
    if !database.New().Health() {
        return c.JSON(500, map[string]string {"message": "KO"})
    }
    return c.JSON(200, map[string]string {"message": "OK"})
}
