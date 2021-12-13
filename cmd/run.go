package cmd

import (
	"github.com/willdavsmith/app/server"
)

func RunServer(port int) {
	server := &server.Server{}
	server.RunWebServer()
}
