package main

import (
	"fmt"
	"os"

	"github.com/aabdelrhmann1/rssagg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
