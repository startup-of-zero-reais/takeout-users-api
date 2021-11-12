package main

import "github.com/startup-of-zero-reais/takeout-users-api/server/infra/server"

func main() {
	s := server.NewServer()
	s.Listen("8081")
}
