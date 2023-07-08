package util

import (
	"fmt"
	"os"
)

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", msg, err)
		os.Exit(1)
	}
}
