package main

// generates misspelling rules for "inital->initial"

import (
	"fmt"
	"strings"
)

var words = `
deinitialization
deinitialize
deinitialized
deinitializes
deinitializing
initial
initialed
initialese
initialing
initialisation
initialisations
initialise
initialised
initialiser
initialisers
initialises
initialising
initialism
initialisms
initializable
initialization
initializations
initialize
initialized
initializer
initializers
initializes
initializing
initialled
initialling
initially
initialness
initials
noninitial
noninitialized
preinitialization
preinitialize
preinitialized
preinitializes
preinitializing
reinitialise
reinitialised
reinitialises
reinitialising
reinitialization
reinitializations
reinitialize
reinitialized
reinitializes
reinitializing
uninitialised
uninitializable
uninitialized
`

func main() {
	for _, word := range strings.Split(strings.TrimSpace(words), "\n") {
		misspell := strings.Replace(word, "initial", "inital", 1)
		fmt.Println(misspell + "->" + word)
	}
}
