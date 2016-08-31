package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const prefix = "WIKITIONARY"

// isAZ returns true if input is only consists of letters a-z (case SENSITIVE)
func isAZ(s string) bool {
	// NO: for _, ch := range s {
	//   This iterates over UTF-8 characters
	slen := len(s)
	for i := 0; i < slen; i++ {
		ch := s[i]
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}

func main() {
	var (
		infile = flag.String("in", "words.txt.gz", "input file")
	)
	flag.Parse()
	wordmap := make(map[string]bool, 2000000)

	fi, err := os.Open(*infile)
	if err != nil {
		log.Fatalf("[%s]: %s", prefix, err)
	}

	fizip, err := gzip.NewReader(fi)
	if err != nil {
		log.Fatalf("[%s]: gzip error %s", prefix, err)
	}

	scanner := bufio.NewScanner(fizip)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(strings.TrimSpace(line), "_")
		for _, word := range words {
			if isAZ(word) {
				wordmap[word] = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading: %s", err)
	}

	wordlist := make([]string, 0, len(wordmap))
	for k, _ := range wordmap {
		wordlist = append(wordlist, k)
	}
	sort.Strings(wordlist)
	for _, word := range wordlist {
		fmt.Println(word)
	}
}
