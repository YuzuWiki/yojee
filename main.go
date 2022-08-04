package main

import (
	"github.com/joho/godotenv"
)

func main() {
	// run service
}

func init() {
	// init some service
	_ = godotenv.Load(".env")
}
