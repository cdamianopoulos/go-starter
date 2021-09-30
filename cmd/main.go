package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/status", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := fmt.Fprint(w, `{"status": "healthy"}`)
		if err != nil {
			fmt.Printf("writing to stream: %v\n", err)
		}
	})
	fmt.Println("Hello, serving on port 3000")
	http.ListenAndServe(":3000", r)
}
