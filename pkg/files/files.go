package files

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jffin/have_ibeen_pwned_multi/pkg/checker"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/errors"
)

func ReadInputFile(fileName string) []string {
	content, err := os.ReadFile(fileName)
	errors.Check("reading input file", err)
	return cleanupInputContent(content)
}

func WriteOutputFile(fileName string, results []checker.Response) {
	data, _ := json.Marshal(results)
	err := os.WriteFile(fileName, data, 0644)
	errors.Check("writing to output file", err)
}

func cleanupInputContent(content []byte) []string {
	var windowsSupportedString string = strings.ReplaceAll(string(content), "\r\n", "\n")
	var withRemovedLastEmptyLine string = strings.TrimRight(windowsSupportedString, "\n")
	return strings.Split(withRemovedLastEmptyLine, "\n")
}
