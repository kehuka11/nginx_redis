package main

import (
	"sample_app2/http/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8081"))
}
