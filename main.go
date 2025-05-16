package main

import "github.com/devinmiller/fem-complete-go/internal/app"

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
}
