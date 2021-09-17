package args

import (
	"flag"
)

const (
	DefaultInputArgs   string = "input"
	DefaultOutputArgs  string = "output"
	DefaultApiKeyArgs  string = "key"
	DefaultApiKeyValue string = "api-key"
	InputFileName      string = "input.txt"
	ResultFileName     string = "results.json"
)

func InitArgs() (inputFile, outputFile, apiKey *string) {
	inputFile = flag.String(DefaultInputArgs, InputFileName, "Input file with targets")
	outputFile = flag.String(DefaultOutputArgs, ResultFileName, "Output file write results to")
	apiKey = flag.String(DefaultApiKeyArgs, DefaultApiKeyValue, "hibp-api-key")
	flag.Parse()

	return
}
