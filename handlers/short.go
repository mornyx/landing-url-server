package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mornyx/landing-url-server/db"
	"github.com/mornyx/landing-url-server/pkg/logx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// URLShortHandler implements business logic for /url/:shortid.
// It writes back the corresponding url based on the shortid parameter.
func URLShortHandler(d db.ShortURLStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortid := c.Param("shortid")
		if shortid == "" {
			c.String(http.StatusBadRequest, "missing path param 'shortid'")
			return
		}
		row, err := d.FindShortURLByShortID(shortid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.String(http.StatusNotFound, "not found")
				return
			}
			logx.Error("failed to get url by shortid", zap.Error(err))
			c.String(http.StatusInternalServerError, "server error")
			return
		}
		c.String(http.StatusOK, row.URL)
	}
}
