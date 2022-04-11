package main

import (
	"github.com/cloverzrg/onefile/api"
	"github.com/cloverzrg/onefile/db"
	"github.com/cloverzrg/onefile/logger"
	"os"
)

func main() {
	var err error
	err = db.Connect()
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	err = api.Start()
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
}
