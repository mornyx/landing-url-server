package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mornyx/landing-url-server/db"
	"github.com/mornyx/landing-url-server/pkg/genid"
	"github.com/mornyx/landing-url-server/pkg/logx"
	"go.uber.org/zap"
)

// URLHandler implements business logic for /url.
// It generates a shortid for the url parameter in the post form and writes is back.
func URLHandler(gen genid.Generator, d db.ShortURLStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.PostForm("url")
		if url == "" {
			c.String(http.StatusBadRequest, "missing form param 'url'")
			return
		}
		for retry := 0; retry < 3; retry++ {
			shortid := gen.Generate(url)
			if err := d.CreateShortURL(&db.ShortURL{URL: url, ShortID: shortid}); err != nil {
				if db.ErrIsSQLiteConstraintUnique(err) {
					time.Sleep(10 * time.Millisecond)
					continue
				}
				logx.Error("failed to insert short url into database", zap.Error(err))
				c.String(http.StatusInternalServerError, "server error")
				return
			}
			c.String(http.StatusOK, shortid)
			return
		}
		logx.Error("shortid conflicted and maximum retry times exceeded")
		c.String(http.StatusInternalServerError, "server error")
	}
}
