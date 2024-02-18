package main

import "fmt"

// App - the struct wjich contains things like
// pointers to database connections
type App struct {
}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")
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
