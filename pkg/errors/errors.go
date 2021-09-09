package errors

import (
	"fmt"
	"os"
)

func Check(errorMessage string, err error) {
	if err != nil {
		fmt.Printf("Error happened during %s: %s\n", errorMessage, err)
		os.Exit(1)
	}
}
