package bar

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hironow/env-gae-run-template/common"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Name string
	Env  string
	Piyo string
}

func Start() {
	var c Config
	err := envconfig.Process("app", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	http.HandleFunc("/", index(&c))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("bar: Defaulting to port %s", port)
	}

	log.Printf("bar: Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func index(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Application:%s\nEnviroment:%s\nMessage:%s\nPiyo:%s", c.Name, c.Env, common.Message(), c.Piyo)
	}
}
