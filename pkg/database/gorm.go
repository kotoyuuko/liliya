package database

import "github.com/jinzhu/gorm"

// Connect returns a database instance
func Connect(dialect, args string) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, args)
	if err != nil {
		return nil, err
	}

	return db, nil
}
