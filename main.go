package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, output io.Writer) error {
	flags := flag.NewFlagSet("training-go-releaser", flag.ContinueOnError)
	showVersion := flags.Bool("version", false, "print version information")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *showVersion {
		fmt.Fprintf(output, "training-go-releaser version %s (commit: %s, date: %s)\n", version, commit, date)
		return nil
	}
	fmt.Fprintln(output, "test 1")
	return nil
}
