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

// Serve creates the http server
func Serve(svc service.Service, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))
	router.Use(middleware.Logger)

	router.Post("/api/post", handlers.MakePostPostHandler(svc))
	router.Get("/api/post/{id}", handlers.MakeGetPostHandler(svc))
	// router.Delete("/api/post/{id}", handlers.MakeDeletePostHandler(svc))

	// TODO: Front end static delivery
	router.Get("/*", handlers.MakeGetHTMLHandler("post.html"))
	router.Get("/create", handlers.MakeGetHTMLHandler("create.html"))
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./front/static"))))

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
