package main

import (
	"time"

	"golang.org/x/time/rate"

	"github.com/jffin/have_ibeen_pwned_multi/pkg/args"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/checker"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/client"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/files"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/structs"
)

func main() {

	inputFile, outputFile, apiKey := args.InitArgs()

	var targetsArray []string = files.ReadInputFile(*inputFile)

	rateLimiter := rate.NewLimiter(rate.Every(1501*time.Millisecond), 1) // 1 request every 1500 milliseconds
	client := client.NewClient(rateLimiter)

	channel := make(chan structs.Response)
	for _, target := range targetsArray {
		go checker.CheckEmail(target, *apiKey, client, channel)
	}

	results := make([]structs.Response, len(targetsArray))
	for index, _ := range results {
		results[index] = <-channel
	}
	files.WriteOutputFile(*outputFile, results)
}
