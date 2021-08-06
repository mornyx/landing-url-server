// Package db wraps the CRUD operations on the database.
// Each domain's data has its own separate XxxStore interface.
// All public interfaces hide the ORM details used internally.
package db

import (
	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ Database = &database{}

// Database aggregates all store interfaces.
type Database interface {
	ShortURLStore

	Migrate() error
}

func NewDatabase(source string) (Database, error) {
	db, err := gorm.Open(sqlite.Open(source))
	if err != nil {
		return nil, err
	}
	return &database{db: db}, nil
}

type database struct {
	db *gorm.DB
}

func (d *database) Migrate() error {
	return d.db.AutoMigrate(&ShortURL{})
}

// ErrIsSQLiteConstraintUnique returns whether err is unique key conflict in SQLite.
func ErrIsSQLiteConstraintUnique(err error) bool {
	serr, ok := err.(sqlite3.Error)
	return ok && serr.Code == sqlite3.ErrConstraint && serr.ExtendedCode == sqlite3.ErrNoExtended(2067)
}
