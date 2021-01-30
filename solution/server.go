package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

func server(app App) {
	listenAddr := flag.String("http.addr", ":8080", "http listen address")
	flag.Parse()
	type responseNumbersGet struct {
		Numbers []int `json:"numbers"`
	}
	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
		URLs := r.URL.Query()["u"]

		response := responseNumbersGet{
			Numbers: app.Collect(r.Context(), URLs),
		}
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
