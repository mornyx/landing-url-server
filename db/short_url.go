package db

import "gorm.io/gorm"

var _ ShortURLStore = &database{}

// ShortURL defines the short_url table schema.
type ShortURL struct {
	gorm.Model
	URL     string `gorm:"column:url"`
	ShortID string `gorm:"column:shortid;uniqueIndex"`
}

func (ShortURL) TableName() string {
	return "short_url"
}

// ShortURLStore provides the CRUD methods for ShortURL without
// exposing any ORM details.
type ShortURLStore interface {
	CreateShortURL(v *ShortURL) error

	FindShortURLByShortID(shortid string) (*ShortURL, error)
}

func (d *database) CreateShortURL(v *ShortURL) error {
	err := d.db.Create(v).Error
	return err
}

func (d *database) FindShortURLByShortID(shortid string) (*ShortURL, error) {
	var v ShortURL
	if err := d.db.Find(&v, "shortid = ?", shortid).Error; err != nil {
		return nil, err
	}
	return &v, nil
}
