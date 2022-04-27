package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/lopz/cs-api-test/internal/database"
	"github.com/lopz/cs-api-test/internal/metric"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes() {

	database.Connect()

	r := chi.NewRouter()

	r.Use(metric.MiddlewarePrometheus)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving request: %s", r.URL.Path)
		host, _ := os.Hostname()
		fmt.Fprintf(w, "Hello, world!\n")
		fmt.Fprintf(w, "Version: 1.0.0\n")
		fmt.Fprintf(w, "Hostname: %s\n", host)
	})

	r.Handle("/metrics", promhttp.Handler())

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/people", func(r chi.Router) {
		r.Post("/", createPerson)
		r.Get("/", getAllPerson)
		r.Route("/{uuid}", func(r chi.Router) {
			r.Get("/", getPerson)       // GET /people/123
			r.Put("/", updatePerson)    // PUT /people/123
			r.Delete("/", deletePerson) // DELETE /people/123
		})

	})

	http.ListenAndServe(":3333", r)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	log.Println("Received Terminate, gracefully shutdown", sig)
}
