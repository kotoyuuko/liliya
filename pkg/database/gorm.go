package database

import (
	"net/url"

	"github.com/jinzhu/gorm"
)

// Connect returns a database instance
func Connect(dialect, args string) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, args)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ArgsString returns a string contains connection information
func ArgsString(host, user, password, name, timezone string) string {
	args := user + ":" + password + "@tcp(" + host + ")/"
	args = args + name + "?charset=utf8&parseTime=True&loc=" + url.QueryEscape(timezone)

	return args
}
