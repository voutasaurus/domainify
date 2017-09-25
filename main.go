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
		log.Fatal("Expected single argument, got: %d", len(os.Args)-1)
	}
	phrase := os.Args[1]
	dd, err := domainify(phrase)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if len(dd) == 0 {
		log.Fatalf("No domains found for: %q", phrase)
	}
	for _, d := range dd {
		fmt.Println(d)
	}
}

type statusError struct {
	got  int
	want int
}

func (s statusError) Error() string {
	return fmt.Sprintf("unexpected status code, want: %d, got: %d", s.want, s.got)
}

func domainify(phrase string) (domains []string, err error) {
	r := strings.NewReplacer(".", "", "*", "")
	phrase = strings.ToLower(phrase)
	res, err := http.Get("https://publicsuffix.org/list/public_suffix_list.dat")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, statusError{got: res.StatusCode, want: 200}
	}
	var dd []string
	s := bufio.NewScanner(res.Body)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 || strings.HasPrefix(line, "//") {
			continue
		}
		line = r.Replace(line)
		if strings.HasSuffix(phrase, line) {
			dd = append(dd, phrase[:len(phrase)-len(line)]+"."+line)
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return dd, nil
}
