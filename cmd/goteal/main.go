package main

import (
	"fmt"
	"os"
)

func main() {
	rootCmd.AddCommand(buildCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
