package server

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/Olyxz16/go-svelte-template/database"
	"github.com/Olyxz16/go-svelte-template/handlers"
)

var (
    host string
    port int
    staticFilepath string
    db database.Service
)


func NewServer() *http.Server {
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
    slog.Info(fmt.Sprintf("Starting server at %s:%d ...", host, port))

    db = database.New()

	return server
}

func init() {
    var err error

    host = os.Getenv("HOST")
    port, err = strconv.Atoi(os.Getenv("PORT"))
    if err != nil {
        panic("Cannot parse port !")
    }

    staticFPFlag := flag.String("staticFilepath", "", "Static files path")
    flag.Parse()
    staticFilepath = *staticFPFlag
    if staticFilepath == "" {
        if os.Getenv("DEBUG") == "true" {
            os.Setenv("API_MODE", "true")
            slog.Warn("API MODE IS ACTIVE. Add --staticFilepath flag to serve static file")
        } else {
            panic("--staticFilepath <path> is not set.")
        }
    } else {
        handlers.StaticFilepath = staticFilepath
    }
}
