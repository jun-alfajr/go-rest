package main

import "github.com/jun-alfajr/go-rest/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":2106"))
}
