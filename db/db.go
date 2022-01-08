package db

import (
	"github.com/akaahmedkamal/go-cli/v1"
	"github.com/akaahmedkamal/go-server/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	db *gorm.DB
}

var shared *Database

func Shared() *Database {
	return shared
}

func New(app *cli.App) *Database {
	cfg := config.Of(app)

	if err := Disconnect(); err != nil {
		panic(err)
	}

	if _, err := Connect(cfg.DbDriver(), cfg.DbUrl()); err != nil {
		panic(err)
	}

	return shared
}

func Connect(driver, url string) (*Database, error) {
	if shared == nil {
		shared = new(Database)
		var err error
		shared.db, err = gorm.Open(sqlite.Open(url), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	return shared, nil
}

func Disconnect() error {
	if shared != nil && shared.db != nil {
		if db, err := shared.db.DB(); err != nil {
			return err
		} else {
			log.Println("[DB]: closing connection...")
			err = db.Close()
			if err == nil {
				log.Println("[DB]: connection closed successfully.")
			}
			return err
		}
	}
	return nil
}
