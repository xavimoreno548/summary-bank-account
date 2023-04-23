package main

import "github.com/xavimoreno548/summary-bank-account/internal/router"

const (
	PORT = "8080"
)

func main() {
	appRouter, err := router.AppRouter()
	if err != nil {
		panic(err)
	}
	err = router.RunApp(PORT, appRouter)
	if err != nil {
		panic(err)
	}
}
