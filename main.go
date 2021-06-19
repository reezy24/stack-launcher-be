package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	PORT = "8090"
	ORIGIN = "http://localhost:3000"
)

func main() {
	r := mux.NewRouter()
	goExecutable, _ := exec.LookPath("go")

	r.HandleFunc("/run", func(rw http.ResponseWriter, r *http.Request) {
		cmdGoVer := &exec.Cmd{
			Path: goExecutable,
			Args: []string{ goExecutable, "version" },
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}
	
		if err := cmdGoVer.Run(); err != nil {
			panic(err)
		}
	})

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{ORIGIN})
	h := handlers.CORS(headers, origins)(r)

	fmt.Printf("Stack launcher listening on port %s.\n", PORT)

	if err := http.ListenAndServe(":" + PORT, h); err != nil {
		panic(err)
	}
}