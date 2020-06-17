package baz

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hironow/env-gae-run-template/common"
)

func Start() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("baz: Defaulting to port %s", port)
	}

	if v := os.Getenv("APP_NAME"); v != "" {
		log.Printf("Application:%s", v)
	}

	log.Printf("baz: Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if v := os.Getenv("APP_ENV"); v != "" {
		fmt.Fprintf(w, "baz: %s on %s\n", common.Message(), v)
	} else {
		fmt.Fprintf(w, "baz: %s\n", common.Message())
	}
}
