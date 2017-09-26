package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/voutasaurus/domainify/domainify"
)

func main() {
	t, err := template.New("list-domains").Parse(htmlListDomains)
	if err != nil {
		log.Fatalf("Could not parse list-domains template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) <= 1 {
			io.WriteString(w, htmlHome)
			return
		}
		dd, err := domainify.Possibilities(r.URL.Path[1:])
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting public suffixes: %v", err), 500)
			return
		}
		if len(dd) == 0 {
			dd = append(dd, "No domains found for "+r.URL.Path[1:])
		}
		if err := t.ExecuteTemplate(w, "list-domains", dd); err != nil {
			http.Error(w, fmt.Sprintf("Error executing template with: %v", dd), 500)
		}
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) <= len("/api/") {
			http.Error(w, "URL path must contain phrase, was empty instead", 400)
			return
		}
		dd, err := domainify.Possibilities(r.URL.Path[len("/api/"):])
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting public suffixes: %v", err), 500)
			return
		}
		if err := json.NewEncoder(w).Encode(struct {
			Domains []string
		}{
			Domains: dd,
		}); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), 500)
		}
	})
	log.Fatal(http.ListenAndServe(":9090", nil))
}
