package main

import (
	"github.com/Olyxz16/go-svelte-template/server"
)

func main() {
    
    server := server.NewServer()
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
	}

}
