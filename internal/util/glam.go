package util

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func GPrint(msg string) error {
	out, err := glamour.Render(msg, "auto")
	if err != nil {
		return fmt.Errorf("glamour render error %w", err)
	}
	fmt.Println(out)
	return nil
}
