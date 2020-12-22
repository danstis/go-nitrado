package main

import (
	"fmt"

	"github.com/danstis/go-nitrado/nitrado"
)

var token string = "YourTokenHere"

func main() {
	api := nitrado.NewClient(token)

	services, _, err := api.Services.List()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Services: %#v", services)
}
