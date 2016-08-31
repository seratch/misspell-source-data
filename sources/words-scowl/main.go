package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

// LoadWordList loads in a list of known-good words
const scowlUS60 = "http://app.aspell.net/create?max_size=60&spelling=US&max_variant=1&diacritic=strip&special=hacker&download=wordlist&encoding=utf-8&format=inline"

func loadWordList(target string) map[string]bool {
	resp, err := http.Get(target)
	if err != nil {
		log.Fatalf("Unable to download: %s", err)
	}
	defer resp.Body.Close()
	out := make(map[string]bool)
	intro := true
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if intro {
			if line == "---" {
				intro = false
			}
			continue
		}

		// we dont use contractions
		if strings.Contains(line, "'") {
			continue
		}

		out[strings.ToLower(line)] = true
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading: %s", err)
	}
	return out
}

func main() {
	outfile := flag.String("out", "words.txt", "Name of output file")
	flag.Parse()

	wordmap := loadWordList(scowlUS60)
	// do any fixups here

	words := make([]string, 0, len(wordmap))
	for k, _ := range wordmap {
		words = append(words, k)
	}
	log.Printf("Got %d words", len(words))

	sort.Strings(words)

	fo, err := os.Create(*outfile)
	if err != nil {
		log.Fatalf("unable to create %s: %s", *outfile, err)
	}

	buf := bufio.NewWriter(fo)
	for _, word := range words {
		buf.WriteString(word + "\n")
	}
	buf.Flush()
	fo.Close()
}
