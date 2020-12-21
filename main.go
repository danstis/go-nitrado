package main

import (
	"log"
)

// Version contains the package version
var Version = "DEVELOPMENT"

// Main entry point for the app.
func main() {
	log.Printf("Version %q", Version)
}
