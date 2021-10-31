package config

import (
	"log"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Message  string
}

func (app *Application) Super() {
	app.Message = "Hello"
}
