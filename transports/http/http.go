package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/blug/config"
	"github.com/sergiosegrera/blug/service"
	"github.com/sergiosegrera/blug/transports/http/handlers"
)

func Serve(svc service.Service, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))
	router.Use(middleware.Logger)

	router.Post("/post", handlers.MakePostPostHandler(svc))
	// router.Get("/post", handlers.MakeGetPostHandler(svc))
	// router.Delete("/post", handlers.MakeDeletePostHandler(svc))

	// TODO: Front end static delivery

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
