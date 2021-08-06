package handlers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mornyx/landing-url-server/db"
	"github.com/stretchr/testify/assert"
)

func TestURLShortHandler_200(t *testing.T) {
	d := db.NewMockDatabase()
	_ = d.CreateShortURL(&db.ShortURL{URL: "abc", ShortID: "def"})
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Params = []gin.Param{{Key: "shortid", Value: "def"}}
	h := URLShortHandler(d)
	h(c)
	assert.Equal(t, http.StatusOK, rw.statusCode)
	assert.Equal(t, []byte("abc"), rw.body)
}

func TestURLShortHandler_400(t *testing.T) {
	d := db.NewMockDatabase()
	_ = d.CreateShortURL(&db.ShortURL{URL: "abc", ShortID: "def"})
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	// c.Params = []gin.Param{{Key: "shortid", Value: "def"}}
	h := URLShortHandler(d)
	h(c)
	assert.Equal(t, http.StatusBadRequest, rw.statusCode)
}

func TestURLShortHandler_404(t *testing.T) {
	d := db.NewMockDatabase()
	// _ = d.CreateShortURL(&db.ShortURL{URL: "abc", ShortID: "def"})
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Params = []gin.Param{{Key: "shortid", Value: "def"}}
	h := URLShortHandler(d)
	h(c)
	assert.Equal(t, http.StatusNotFound, rw.statusCode)
}

func TestURLShortHandler_500(t *testing.T) {
	d := db.NewMockDatabase()
	d.SetReturnedError(errors.New("mock error"))
	_ = d.CreateShortURL(&db.ShortURL{URL: "abc", ShortID: "def"})
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Params = []gin.Param{{Key: "shortid", Value: "def"}}
	h := URLShortHandler(d)
	h(c)
	assert.Equal(t, http.StatusInternalServerError, rw.statusCode)
	assert.Equal(t, []byte("server error"), rw.body)
}
