package main

import (
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

var checkingExtensions int

var path string
var format string
var revision string

var committer bool

// extensionFLAG

func firstCheck() {
	if committer {
		author = "committer"

	}
	if len(flagExtensions) > 0 {
		checkingExtensions |= 1
		for _, item := range flagExtensions {
			extensions[item] |= 1
		}
	}
	if len(languages) > 0 {
		checkingExtensions |= 2
		for _, i := range languages {
			for _, j := range LangExtension[i] {
				extensions[j] |= 2
			}
		}
	}
}

func run(app *cobra.Command, args []string) {
	firstCheck()
	ls, _ := GitLsTree(revision)
	listFiles := strings.Split(string(ls), "\n")
	for _, file := range listFiles {
		if len(file) == 0 {
			continue
		}
		counter := 0
		start := 0
		for i, j := range file {
			if unicode.IsSpace(j) {
				counter++
			} else if counter >= 3 {
				start = i
				break
			}
		}
		thisFile := file[start:]
		if GeneralCheck(thisFile) {
			calcStats(thisFile)
		}
	}
	sortByElem()
	printStats()
}
