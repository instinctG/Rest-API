package main

import (
	coment "github.com/instinctG/Rest-API/internal/comment"
	"github.com/instinctG/Rest-API/internal/config"
	http2 "github.com/instinctG/Rest-API/internal/transport/http"
	log "github.com/sirupsen/logrus"

	"github.com/instinctG/Rest-API/internal/database"
	"net/http"
)

// App - the struct which contains information about our app
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"AppName":    app.Name,
		"AppVersion": app.Version,
	}).Info("Setting Up Our App")
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDataBase(database.Config{
		Host:     cfg.Config.Host,
		Port:     cfg.Config.Port,
		User:     cfg.Config.User,
		DBName:   cfg.Config.DBName,
		Password: cfg.Config.Password,
		SSLMode:  cfg.Config.SSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	commentService := coment.NewService(db)

	handler := http2.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

// our main entrypoint for the application
func main() {
	app := App{
		Name:    "Comment API",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error Starting Up")
		log.Fatal(err)
	}
}
