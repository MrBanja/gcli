package util

import (
	"fmt"
	"github.com/charmbracelet/glamour"
)

func PrintOrExit(msg string) {
	out, err := glamour.Render(msg, "auto")
	HandleError(err, "Glamour render error")
	fmt.Println(out)
}
