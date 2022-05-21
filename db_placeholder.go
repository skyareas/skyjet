//go:build !gorm
// +build !gorm

package skyjet

type DbClient struct {
}

var db *DbClient

func (c *DbClient) Disconnect() error {
	return nil
}
