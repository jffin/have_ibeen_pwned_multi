package errors

import (
	"errors"
	"testing"
)

const (
	ExpectedExitCode    int    = 1
	DefaultErrorMessage string = "some error"
)

func TestCheck(t *testing.T) {
	// Save current function and restore at the end:
	oldOsExit := OsExit
	defer func() { OsExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	OsExit = myExit

	Check(DefaultErrorMessage, errors.New(DefaultErrorMessage))

	if exp := ExpectedExitCode; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}

	Check("", nil)
}
