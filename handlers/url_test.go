package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mornyx/landing-url-server/db"
	"github.com/mornyx/landing-url-server/pkg/genid"
	"github.com/stretchr/testify/assert"
)

func TestURLHandler_200(t *testing.T) {
	d := db.NewMockDatabase()
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Request = &http.Request{PostForm: url.Values{"url": []string{"abc"}}}
	h := URLHandler(genid.MustNewGenerator(), d)
	h(c)
	assert.Equal(t, http.StatusOK, rw.statusCode)
	assert.NotEmpty(t, rw.body)
}

func TestURLHandler_400(t *testing.T) {
	d := db.NewMockDatabase()
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Request = &http.Request{PostForm: url.Values{}}
	h := URLHandler(genid.MustNewGenerator(), d)
	h(c)
	assert.Equal(t, http.StatusBadRequest, rw.statusCode)
}

func TestURLHandler_500_insert(t *testing.T) {
	d := db.NewMockDatabase()
	d.SetReturnedError(errors.New("mock error"))
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Request = &http.Request{PostForm: url.Values{"url": []string{"abc"}}}
	h := URLHandler(genid.MustNewGenerator(), d)
	h(c)
	assert.Equal(t, http.StatusInternalServerError, rw.statusCode)
}

func TestURLHandler_500_conflict(t *testing.T) {
	d := db.NewMockDatabase()
	_ = d.CreateShortURL(&db.ShortURL{URL: "", ShortID: "CONFLICT"})
	rw := newMockResponseWriter()
	c, _ := gin.CreateTestContext(rw)
	c.Request = &http.Request{PostForm: url.Values{"url": []string{"abc"}}}
	h := URLHandler(&mockConflictGenerator{}, d)
	h(c)
	assert.Equal(t, http.StatusInternalServerError, rw.statusCode)
}

var _ genid.Generator = &mockConflictGenerator{}

type mockConflictGenerator struct{}

func (m *mockConflictGenerator) Generate(k string) string {
	return "CONFLICT"
}
