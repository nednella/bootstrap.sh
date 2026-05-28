package main

import (
	"os"

	"github.com/nednella/bootstrap.sh/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
