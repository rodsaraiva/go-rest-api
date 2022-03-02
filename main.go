package main

import (
	"github.com/rodsaraiva/go-rest-api/app"
	"github.com/rodsaraiva/go-rest-api/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()

}
