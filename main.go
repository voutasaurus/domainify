package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 || len(os.Args) > 3 {
		log.Fatalf("Expected single argument, got: %d", len(os.Args)-1)
	}
	dd, err := domainify(strings.ToLower(os.Args[1]))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if len(dd) == 0 {
		log.Fatalf("No domains found for: %q", os.Args[1])
	}
	for _, d := range dd {
		fmt.Println(d)
	}
}

func domainify(phrase string) (domains []string, err error) {
	res, err := http.Get("https://publicsuffix.org/list/public_suffix_list.dat")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("want status: %d, got: %d", 200, res.StatusCode)
	}
	s := bufio.NewScanner(res.Body)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 || strings.HasPrefix(line, "//") {
			continue
		}
		suffix := strings.NewReplacer(".", "", "*", "").Replace(line)
		if strings.HasSuffix(phrase, suffix) {
			if strings.HasPrefix(line, "*.") {
				line = line[2:]
			}
			domains = append(domains, phrase[:len(phrase)-len(suffix)]+"."+line)
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}
