package main

import "github.com/xavimoreno548/summary-bank-account/internal/router"

const (
	PORT = "8080"
)

func main() {
	err := router.RunApp(PORT, router.AppRouter())
	if err != nil {
		panic(err)
	}
}
