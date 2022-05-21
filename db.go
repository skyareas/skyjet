//go:build gorm
// +build gorm

package skyjet

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// supported database drivers.
const (
	DriverSQLite  = "sqlite"
	DriverSQLite3 = "sqlite3"
)

type DbClient struct {
	gorm.DB
}

// db instance of the database.
var db *DbClient

// DB returns the shared/global instance of the
// database if found, otherwise, it initializes
// a new shared instance and returns it.
func DB(opts ...gorm.Option) *DbClient {
	if db == nil {
		driver := app.cfg.Db.Driver
		url := app.cfg.Db.Url

		var dial gorm.Dialector

		switch driver {
		case DriverSQLite:
		case DriverSQLite3:
			dial = sqlite.Open(url)
		default:
			app.log.Fatalf("Unsupported database driver \"%s\"\n", driver)
		}

		client, err := gorm.Open(dial, opts...)
		if err != nil {
			return nil
		}

		db = &DbClient{*client}
	}

	return db
}

// Disconnect closes the database connection if open.
func (c *DbClient) Disconnect() error {
	d, err := c.DB.DB()
	if err != nil {
		return err
	}
	err = d.Close()
	if err != nil {
		app.log.Printf("database connection closed with error: %s", err.Error())
	} else {
		app.log.Println("database connection closed successfully")
	}
	return err
}
