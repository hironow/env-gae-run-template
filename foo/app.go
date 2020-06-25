package foo

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hironow/env-gae-run-template/common"
)

func Start() {
	cfg, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", cfg)

	http.HandleFunc("/", index(cfg))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("foo: Defaulting to port %s", port)
	}

	log.Printf("foo: Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func index(c *common.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Application: %s\nEnviroment: %s\nMessage: %s\nHoge: %s\n", c.App.Name, c.App.Env, common.Message(), c.App.Hoge)

		z, err := common.GetZone(c)
		if err != nil {
			fmt.Fprintf(w, "Zone Error: %s\n", err.Error())
		}
		fmt.Fprintf(w, "Zone: %s\n", z)

		p, err := common.GetProjectID(c)
		if err != nil {
			fmt.Fprintf(w, "ProjectID Error: %s\n", err.Error())
		}
		fmt.Fprintf(w, "ProjectID: %s\n", p)

		fmt.Fprintf(w, "Config: %+v\n", c)

		hn := common.GetHostName(c)
		if err != nil {
			fmt.Fprintf(w, "HostName Error: %s\n", err.Error())
		}
		fmt.Fprintf(w, "HostName: %s\n", hn)

	}
}
