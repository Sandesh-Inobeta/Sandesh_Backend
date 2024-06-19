package main

import (
	"fmt"
	"log"
	"net/http"
	"sandesh/Redis"
	router "sandesh/routers"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	port := app.Config.Port
	fmt.Printf("Serving app on port %s", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}
	return srv.ListenAndServe()
}

func main() {
	Redis.GetClient()
	defer Redis.CloseClient()

	config := Config{
		Port: "3000",
	}

	app := &Application{
		Config: config,
	}

	err := app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
