package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := fmt.Fprint(w, `{"status": "healthy"}`)
		if err != nil {
			fmt.Printf("writing to stream: %v\n", err)
		}
	})

	fmt.Println("Hello, serving on port 3004")
	http.ListenAndServe(":3004", r)
}
