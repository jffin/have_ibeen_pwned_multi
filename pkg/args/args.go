package args

import (
	"flag"
)

const (
	DEFAULT_INPUT_ARGS    string = "input"
	DEFAULT_OUTPUT_ARGS   string = "output"
	DEFAULT_API_KEY_ARGS  string = "key"
	DEFAULT_API_KEY_VALUE string = "api-key"
	INPUT_FILE_NAME       string = "input.txt"
	RESULT_FILE_NAME      string = "results.json"
)

func InitArgs() (inputFile, outputFile, apiKey *string) {
	inputFile = flag.String(DEFAULT_INPUT_ARGS, INPUT_FILE_NAME, "input file with targets")
	outputFile = flag.String(DEFAULT_OUTPUT_ARGS, RESULT_FILE_NAME, "output file write results to")
	apiKey = flag.String(DEFAULT_API_KEY_ARGS, DEFAULT_API_KEY_VALUE, "hibp-api-key")
	flag.Parse()

	return
}
