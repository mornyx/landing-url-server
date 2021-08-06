package db

import (
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

var _ Database = &MockDatabase{}

// MockDatabase implements a fake database for testing purposes.
type MockDatabase struct {
	m   map[string]*ShortURL
	err error
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		m: map[string]*ShortURL{},
	}
}

func (d *MockDatabase) CreateShortURL(v *ShortURL) error {
	if d.err != nil {
		return d.err
	}
	if _, ok := d.m[v.ShortID]; ok {
		return sqlite3.Error{
			Code:         sqlite3.ErrConstraint,
			ExtendedCode: sqlite3.ErrNoExtended(2067),
		}
	}
	d.m[v.ShortID] = v
	return nil
}

func (d *MockDatabase) FindShortURLByShortID(shortid string) (*ShortURL, error) {
	if d.err != nil {
		return nil, d.err
	}
	if v, ok := d.m[shortid]; ok {
		return v, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (d *MockDatabase) Migrate() error {
	return d.err
}

/* for mocking */

// SetReturnedError fill in a fake err which is returned when another method is called.
func (d *MockDatabase) SetReturnedError(err error) {
	d.err = err
}
