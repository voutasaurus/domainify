package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voutasaurus/domainify/domainify"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) <= 1 {
			http.Error(w, "URL path must contain phrase, was empty instead", 400)
		}
		dd, err := domainify.Possibilities(r.URL.Path[1:])
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting public suffixes: %v", err), 500)
		}
		if err := json.NewEncoder(w).Encode(struct {
			Domains []string
		}{
			Domains: dd,
		}); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), 500)
		}
	})
	http.ListenAndServe(":9090", nil)
}
