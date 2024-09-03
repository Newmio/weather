package main

import "weather/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
