package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockDatabase(t *testing.T) {
	db := NewMockDatabase()
	err := db.Migrate()
	assert.NoError(t, err)
	err = db.CreateShortURL(&ShortURL{
		URL:     "abc",
		ShortID: "def",
	})
	assert.NoError(t, err)
	err = db.CreateShortURL(&ShortURL{
		URL:     "abc",
		ShortID: "def",
	})
	assert.Error(t, err)
	assert.True(t, ErrIsSQLiteConstraintUnique(err))
	v, err := db.FindShortURLByShortID("def")
	assert.NoError(t, err)
	assert.Equal(t, "abc", v.URL)
}
