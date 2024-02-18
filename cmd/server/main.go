package main

import (
	"fmt"
	"github.com/instinctG/Rest-API/Internal/transport"
	"net/http"
)

// App - the struct wjich contains things like
// pointers to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	handler := transport.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

// our main entrypoint for the application
func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}
