package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmc/goteal/build"
)

func init() {
	buildCmd.Flags().BoolP("debug", "d", false, "if true, prints additional debugging output")
	buildCmd.Flags().Bool("dump-ssa", false, "if true, dump the SSA representation of the program")
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build compiles a go package into AVM/TEAL bytecode",
	Run: func(cmd *cobra.Command, args []string) {
		b := build.New()
		if err := b.LoadSources(args...); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		b.Debug = cmd.Flags().Lookup("debug").Value.String() == "true"
		b.DumpSSA = cmd.Flags().Lookup("dump-ssa").Value.String() == "true"
		prg, err := b.Build()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// render TEAL output
		fmt.Println(prg)
	},
}
