package main

import (
	"fmt"
	"log"
	"os"

	"github.com/voutasaurus/domainify/domainify"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 || len(os.Args) > 3 {
		log.Fatalf("Expected single argument, got: %d", len(os.Args)-1)
	}
	dd, err := domainify.Possibilities(os.Args[1])
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
