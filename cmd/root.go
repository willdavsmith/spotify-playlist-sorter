package cmd

import (
	"flag"
)

func Execute() {
	port := flag.Int("port", 8080, "Port to listen to")

	flag.Parse()

	RunServer(*port)
}
