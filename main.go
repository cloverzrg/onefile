package main

import (
	"github.com/cloverzrg/onefile/api"
	"github.com/cloverzrg/onefile/db"
	"github.com/cloverzrg/onefile/logger"
	"github.com/cloverzrg/onefile/model"
	"os"
)

func main() {
	var err error
	err = api.Start()
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
}

func init() {
	var err error
	err = db.Connect()
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	err = db.DB.AutoMigrate(&model.Token{})
	if err != nil {
		logger.Panic(err)
	}
}
