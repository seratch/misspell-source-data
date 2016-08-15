package main

import (
	"flag"
	"log"
)

// returns true if any character is repeated more than N times
func repeated(s string, n int) bool {
	slen := len(s)
	if slen < n {
		return false
	}
	ch := s[0]
	count := 1
	for i := 1; i < slen; i++ {
		cnext := s[i]
		if cnext != ch {
			ch = cnext
			count = 1
			continue
		}
		count++
		if count == n {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("Need an arg!")
	}

	for _, arg := range args {
		log.Printf("%s: %v", arg, repeated(arg, 4))
	}
}
