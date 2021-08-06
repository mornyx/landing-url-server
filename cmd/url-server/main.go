package main

import (
	"flag"
	"log"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mornyx/landing-url-server/db"
	"github.com/mornyx/landing-url-server/handlers"
	"github.com/mornyx/landing-url-server/middlewares"
	"github.com/mornyx/landing-url-server/pkg/genid"
	"github.com/mornyx/landing-url-server/pkg/logx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	flagDebugMode     = flag.Bool("debug-mode", false, "enable gin debug mode")
	flagDBPath        = flag.String("db-path", ":memory:", "sqlite file path")
	flagListen        = flag.String("listen", ":8080", "<IP:PORT> address to listen")
	flagLogLevel      = zap.LevelFlag("log-level", zap.InfoLevel, "zap log level")
)

func init() {
	flag.Parse()
	if *flagDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	// Prepare database.
	d, err := db.NewDatabase(*flagDBPath)
	if err != nil {
		logx.Fatal("failed to open database", zap.String("path", *flagDBPath))
	}
	if err := d.Migrate(); err != nil {
		logx.Fatal("failed to migrate database", zap.String("path", *flagDBPath))
	}

	// Bind middlewares and routes.
	r := gin.New()
	r.Use(ginzap.Ginzap(logx.Logger(), time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logx.Logger(), true))
	r.Use(middlewares.MetricsMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.POST("/url", handlers.URLHandler(genid.MustNewGenerator(), d))
	r.GET("/url/:shortid", handlers.URLShortHandler(d))

	// Listen and serve.
	if err := r.Run(*flagListen); err != nil {
		log.Fatal("failed to run server", zap.Error(err))
	}
}
