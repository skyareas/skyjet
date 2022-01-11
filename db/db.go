package db

import (
	"github.com/akaahmedkamal/go-server/app"
	"github.com/akaahmedkamal/go-server/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// shared instance of the database.
var shared *gorm.DB

// Shared returns the shared/global instance of the
// database if found, otherwise, it initializes
// a new shared instance and returns it.
func Shared(opts ...gorm.Option) (*gorm.DB, error) {
	if shared == nil {
		driver := config.Shared().Db.Driver
		url := config.Shared().Db.Url

		var dial gorm.Dialector

		switch driver {
		case DriverSQLite:
		case DriverSQLite3:
			dial = sqlite.Open(url)
			break
		default:
			app.Shared().Log().Fatalf("[DB]: Unsupported driver \"%s\"\n", driver)
		}

		var err error
		if shared, err = gorm.Open(dial, opts...); err != nil {
			return nil, err
		}
	}

	return shared, nil
}

// Disconnect closes the shared/global database
// connection if found.
func Disconnect() error {
	if shared != nil {
		db, err := shared.DB()
		if err != nil {
			return err
		}
		app.Shared().Log().Infoln("[DB]: closing connection...")
		err = db.Close()
		if err == nil {
			app.Shared().Log().Infoln("[DB]: connection closed successfully.")
		}
		return err
	}
	return nil
}
