package main

import (
	"sample_app1/http/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
