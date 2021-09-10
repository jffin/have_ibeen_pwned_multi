package errors

import (
	"fmt"
	"os"
)

var OsExit = os.Exit

func Check(errorMessage string, err error) {
	if err != nil {
		fmt.Printf("Error happened during %s: %s\n", errorMessage, err)
		OsExit(1)
	}
}
