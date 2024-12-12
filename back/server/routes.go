package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Olyxz16/go-svelte-template/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Static("/", staticFilepath)
    
    /* Static pages */
    e.GET("/*", handlers.Index)
    
    /* Api endpoints */

    /* Health checks */
    e.GET("/health", handlers.Health)

	return e
}
