package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	PORT = "8090"
	ORIGIN = "http://localhost:3000"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("ohai")
	})

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{ORIGIN})
	h := handlers.CORS(headers, origins)(r)

	fmt.Printf("Stack launcher listening on port %s.\n", PORT)

	if err := http.ListenAndServe(":" + PORT, h); err != nil {
		panic(err)
	}
}