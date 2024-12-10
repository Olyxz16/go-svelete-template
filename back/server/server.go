package server

import (
	"flag"
	"fmt"
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
)

type Server struct {
	port int 
	db database.Service
}

func NewServer() *http.Server {
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

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
        panic("Missing static files path !") 
    }
    handlers.StaticFilepath = staticFilepath
}
