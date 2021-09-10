package args

import (
	"flag"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/constants"
)

func InitArgs() (inputFile, outputFile, apiKey *string) {
	inputFile = flag.String("input", constants.INPUT_FILE_NAME, "input file with targets")
	outputFile = flag.String("output", constants.RESULT_FILE_NAME, "output file write results to")
	apiKey = flag.String("key", "api-key", "hibp-api-key")
	flag.Parse()

	return
}
