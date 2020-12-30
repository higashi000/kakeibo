package main

import "github.com/higashi000/kakeibo/router"

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":5000"))
}
