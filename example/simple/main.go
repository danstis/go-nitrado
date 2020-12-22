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

	fmt.Printf("Services: %#v", services)
}
