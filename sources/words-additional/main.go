package main

// words that not in aspell dictionary (yet)
// likely to be technical in computer related fields

// may or may not will to use for clustering
// (i.e. maybe just completely ignore)
//

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var words = []string{
	"subscripting", // https://github.com/client9/misspell-source-data/issues/6
	"unprojected",  // https://en.wiktionary.org/wiki/unprojected
	"updaters",
	"templatize", // https://en.wiktionary.org/wiki/templatize
	"requesters",
	"requestors", // alternate spelling
	"parallelise",
	"parallelize",
	"hammers",
	"convertors",       // real word
	"implementational", // real word
	"perceptron",       // real word
	"perceptrons",      // real word
	"intensional",      // real word, http://foldoc.org/intensional
	"exploder",         //real word
	"exploders",        // real word
	"duplicative",      // real word
	"computable",       // real word
	"incomputable",     // real word
	"responders",
	"parenthesised", // word
	"deleters",
	"releasers",
	"unindented",
	"upgraders",
	"cumulate",
	"positionals",
	"bundlers",
	"bundler",
	"misqualified",
	"falsey",    // technical word for "false value type"
	"bogons",    // technical word
	"bogon",     // technical word
	"expandos",  // "technical" word, for something that expands (see jQuery)
	"accessors", // technical word
	"accessor",  // technical word
	"postless",
	"codifications",
	"cleaner",     // real word
	"thirty",      // real word
	"expresssion", // real word
	"parens",      // common for parenthesis
	"recognizers",
	"administra", // Spanish?
	"relativize",
	"dependee",
	"functionals",
	"benchmarkers",
}

func main() {
	outfile := flag.String("out", "words.txt", "Name of output file")
	flag.Parse()
	sort.Strings(words)
	for _, word := range words {
		fmt.Println(word)
	}
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
