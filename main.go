package main

import (
	"os"

	"github.com/deantech-io/noaa-indexer/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
