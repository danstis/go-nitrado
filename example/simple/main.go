// Package main provides a simple example of using the go-nitrado library.
package main

import (
	"fmt"
	"os"

	"github.com/danstis/go-nitrado/nitrado"
)

var token string = os.Getenv("nitradoToken")

func main() {
	api := nitrado.NewClient(token)

	services, _, err := api.Services.List()
	if err != nil {
		panic(err)
	}

	for _, s := range *services {
		gs, _, _ := api.GameServers.Get(s.ID)
		fmt.Printf("GameServer for %q: %q\n", s.Details.Name, gs.GameHuman)
	}
}
