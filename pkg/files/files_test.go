package files

import (
	"fmt"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/checker"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/constants"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/errors"
	"os"
	"testing"
)

const (
	EXPECTED_EXIT_CODE      int    = 1
	EXPECTED_CONTENT_LENGTH int    = 4
	DEFAULT_FILES_PATH      string = "../../files/"
)

func TestReadInputFile(t *testing.T) {
	// Save current function and restore at the end:
	oldOsExit := errors.OsExit
	defer func() { errors.OsExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	errors.OsExit = myExit

	ReadInputFile(constants.INPUT_FILE_NAME)

	if exp := EXPECTED_EXIT_CODE; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}

	inputFile := fmt.Sprintf("%s/%s", DEFAULT_FILES_PATH, constants.INPUT_FILE_NAME)
	content := ReadInputFile(inputFile)
	if len(content) != EXPECTED_CONTENT_LENGTH {
		t.Errorf("Expected content length: %d, got: %d", EXPECTED_CONTENT_LENGTH, len(content))
	}
}

func TestWriteOutputFile(t *testing.T) {
	var test_results []checker.Response = []checker.Response{
		checker.Response{
			Target: "email@me.com",
			Data:   nil,
		},
	}

	outputFile := fmt.Sprintf("%s/%s", DEFAULT_FILES_PATH, constants.RESULT_FILE_NAME)
	if err := os.Remove(outputFile); err != nil {
		t.Errorf("Can't remove output file")
	}
	WriteOutputFile(outputFile, test_results)
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Expected output file does not found")
	}
}
