package files

import (
	"fmt"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/args"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/checker"
	"github.com/jffin/have_ibeen_pwned_multi/pkg/errors"
	"os"
	"testing"
)

const (
	ExpectedExitCode      int    = 1
	ExpectedContentLength int    = 4
	DefaultFilesPath      string = "../../files/"
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

	ReadInputFile(args.InputFileName)

	if exp := ExpectedExitCode; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}

	inputFile := fmt.Sprintf("%s/%s", DefaultFilesPath, args.InputFileName)
	content := ReadInputFile(inputFile)
	if len(content) != ExpectedContentLength {
		t.Errorf("Expected content length: %d, got: %d", ExpectedContentLength, len(content))
	}
}

func TestWriteOutputFile(t *testing.T) {
	var testResults = []checker.Response{
		{
			Target: "email@me.com",
			Data:   nil,
		},
	}

	outputFile := fmt.Sprintf("%s/%s", DefaultFilesPath, args.ResultFileName)
	if err := os.Remove(outputFile); err != nil {
		t.Errorf("Can't remove output file")
	}
	WriteOutputFile(outputFile, testResults)
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Expected output file does not found")
	}
}

func TestCleanupInputContent(t *testing.T) {
	inputFile := fmt.Sprintf("%s/%s", DefaultFilesPath, args.InputFileName)
	content, _ := os.ReadFile(inputFile)

	var output = cleanupInputContent(content)
	if len(output) != ExpectedContentLength {
		t.Errorf("Expected output length: %d, got %d", ExpectedContentLength, len(content))
	}
}
