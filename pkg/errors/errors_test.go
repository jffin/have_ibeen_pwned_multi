package errors

import (
	"errors"
	"testing"
)

const (
	EXPECTED_EXIT_CODE    int    = 1
	DEFAULT_ERROR_MESSAGE string = "some error"
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

	Check(DEFAULT_ERROR_MESSAGE, errors.New(DEFAULT_ERROR_MESSAGE))

	if exp := EXPECTED_EXIT_CODE; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}

	Check("", nil)
}
