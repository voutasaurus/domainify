package domainify

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func Possibilities(phrase string) (domains []string, err error) {
	res, err := http.Get("https://publicsuffix.org/list/public_suffix_list.dat")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("want status: %d, got: %d", 200, res.StatusCode)
	}
	phrase = strings.ToLower(phrase)
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
