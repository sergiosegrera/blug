package main

import (
	"os"
	"os/signal"

	"github.com/sergiosegrera/blug/config"
	"github.com/sergiosegrera/blug/db/mockdb"
	"github.com/sergiosegrera/blug/models"
	"github.com/sergiosegrera/blug/service"
	"github.com/sergiosegrera/blug/transports/http"
	"go.uber.org/zap"
)

func main() {
	// Load env
	conf := config.New()

	// Start logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Connect to db
	//	db, err := redisdb.New(conf)
	db, err := mockdb.New(conf)
	db.CreatePost(&models.Post{
		ID:       0,
		Title:    "Lea",
		HTML:     "<h1>Allo Lea</h1>",
		Markdown: "# Allo Lea",
	})
	if err != nil {
		logger.Fatal("error connecting to db")
	}

	// Create service
	shortService := &service.BlugService{
		DB:     db,
		Logger: logger,
	}

	// Start http server
	go func() {
		logger.Info("starting the http server", zap.String("port", "8080"))
		err := http.Serve(shortService, conf)
		if err != nil {
			logger.Error("http server panic", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info("exited", zap.String("sig", sig.String()))
}
