package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/YuzuWiki/yojee/web"
)

func main() {
	// run service
	web.Start(27100)

}

func init() {
	// init some service
	_ = godotenv.Load(".env")

}
