package utils

import (
	"flag"
)

func Flagos() *string {
	port := flag.String("port", "8081", "Port to run the server")
	flag.Parse()
	return port
}
