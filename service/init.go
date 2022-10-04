package service

import (
	"os"

	"github.com/YuzuWiki/yojee/service/web"
)

func Start() {
	web.Start(os.Getenv("LISTEN_PORT"))
}
