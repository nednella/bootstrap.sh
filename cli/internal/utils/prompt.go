package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/nednella/bootstrap.sh/internal/ui"
)

// return true on a y/yes answer
func Confirm(prompt string) bool {
	ui.Warn(prompt + " [y/N]")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return false
	}
	answer := strings.ToLower(strings.TrimSpace(line))
	return answer == "y" || answer == "yes"
}
