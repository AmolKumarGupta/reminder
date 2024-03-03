package util

import "fmt"

func Bold(text string) string {
	return fmt.Sprintf("\x1b[1;39m%s\x1b[0m", text)
}
