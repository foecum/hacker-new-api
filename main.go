package main

import (
	"github.com/foecum/hackernews/item"

	"github.com/foecum/going/server"

	"github.com/foecum/hackernews/user"
)

func main() {
	srv := server.New("", "8080")

	var routes []server.Routes
	routes = append(routes, item.Item{})
	routes = append(routes, user.User{})

	srv.RegisterAllRoutes(routes)
	srv.StartServer()
}
