package db

type Database struct {
}

var shared *Database

func Shared() *Database {
	return shared
}

func Connect(driver, url string) (*Database, error) {
	return shared, nil
}

func Disconnect() error {
	return nil
}
