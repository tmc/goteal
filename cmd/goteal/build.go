package main

import (
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build compiles a go package into AVM/TEAL bytecode",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
