package main

import (
	"os"

	"github.com/CCOLLOT/coop-app-1/cmd"
)

func main() {
	if err := cmd.InitAndRunCommand(); err != nil {
		os.Exit(3)
	}
}
