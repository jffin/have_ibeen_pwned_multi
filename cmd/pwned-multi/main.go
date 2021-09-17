package main

import (
	"github.com/jffin/have_ibeen_pwned_multi/pkg/args"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/checker"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/files"
)

func main() {

	inputFile, outputFile, apiKey := args.InitArgs()

	var targetsArray = files.ReadInputFile(*inputFile)
	results := checker.StartCheck(targetsArray, *apiKey)
	files.WriteOutputFile(*outputFile, results)
}
