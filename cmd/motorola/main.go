package main

import "github.com/TheGreaterHeptavirate/motorola/pkg/app"

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		panic(err)
	}
}
