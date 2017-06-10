package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func addOrPanic(dict map[string]string, key, value string) {
	if val, ok := dict[key]; ok && val != value {
		log.Printf("Removing duplicate mismatch: %q -> %q or %q", key, val, value)
		delete(dict, key)
		return
	}

	// this happens for captialization rules
	//
	// english->English
	//
	// variations will generate
	//
	// engish->English, English->English, ENGLISH->ENGLISH
	//
	// so we ignore them
	if key == value {
		return
	}

	dict[key] = value
}

func mergeDict(a, b map[string]string) {
	for k, v := range b {
		addOrPanic(a, k, v)
	}
}

func parseWikipediaFormat(text string) map[string]string {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	dict := make(map[string]string, len(lines))
	for _, line := range lines {
		line = strings.ToLower(strings.TrimSpace(line))
		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			log.Fatalf(fmt.Sprintf("failed parsing %q", line))
		}
		spellings := strings.Split(parts[1], ",")
		dict[parts[0]] = strings.TrimSpace(spellings[0])
	}
	return dict
}

type sortByLen []string

func (a sortByLen) Len() int      { return len(a) }
func (a sortByLen) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByLen) Less(i, j int) bool {
	if len(a[i]) == len(a[j]) {
		return a[i] < a[j]
	}
	// INVERTED
	return len(a[i]) > len(a[j])
}

func main() {
	out := flag.String("o", "words.go", "output file")
	flag.Parse()
	fo, err := os.Create(*out)
	if err != nil {
		log.Fatalf("unable to write %s: %s", *out, err)
	}
	defer fo.Close()
	fo.WriteString("package misspell\n\n")

	// standard for machine generated files
	// https://github.com/golang/go/issues/13560
	// https://godoc.org/github.com/shurcooL/go/generated
	// `^// Code generated .* DO NOT EDIT\.$`
	fo.WriteString("// Code generated automatically.  DO NOT EDIT.\n\n")

	// create main word list
	dict := make(map[string]string)
	mergeDict(dict, dictWikipedia())
	mergeDict(dict, dictAdditions())
	mergeDict(dict, dictReddit())
	log.Printf("With Reddit: %d", len(dict))
	//mergeDict(dict, dictWikipediaArticles())
	//log.Printf("With Wikipedia Articles: %d", len(dict))

	words := make([]string, 0, len(dict))
	for k := range dict {
		words = append(words, k)
	}
	sort.Sort(sortByLen(words))
	fo.WriteString("// DictMain is the main rule set, not including locale-specific spellings\n")
	fo.WriteString("var DictMain = []string{\n")
	for _, word := range words {
		fo.WriteString(fmt.Sprintf("\t%q, %q,\n", word, dict[word]))
	}
	fo.WriteString("}\n\n")

	dict = make(map[string]string)
	mergeDict(dict, dictAmerican())
	words = make([]string, 0, len(dict))
	for k := range dict {
		words = append(words, k)
	}
	sort.Sort(sortByLen(words))
	fo.WriteString("// DictAmerican converts UK spellings to US spellings\n")
	fo.WriteString("var DictAmerican = []string{\n")
	for _, word := range words {
		fo.WriteString(fmt.Sprintf("\t%q, %q,\n", word, dict[word]))
	}
	fo.WriteString("}\n\n")

	dict = make(map[string]string)
	mergeDict(dict, dictBritish())
	words = make([]string, 0, len(dict))
	for k := range dict {
		words = append(words, k)
	}
	sort.Sort(sortByLen(words))
	fo.WriteString("// DictBritish converts US spellings to UK spellings\n")
	fo.WriteString("var DictBritish = []string{\n")
	for _, word := range words {
		fo.WriteString(fmt.Sprintf("\t%q, %q,\n", word, dict[word]))
	}
	fo.WriteString("}\n")
}
