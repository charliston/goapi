package main

import (
	"github.com/charliston/goapi/config"
	"github.com/charliston/goapi/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// INITIALIZE ROUTES
	router.Init()

}
