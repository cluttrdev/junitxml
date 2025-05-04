package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/cluttrdev/cli"
)

func main() {
	if err := exec(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func exec(ctx context.Context) error {
	cmd := configure()

	args := os.Args[1:]
	opts := []cli.ParseOption{}

	if err := cmd.Parse(args, opts...); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		} else {
			return fmt.Errorf("parse arguments: %w", err)
		}
	}

	return cmd.Run(ctx)
}

func configure() *cli.Command {
	return &cli.Command{
		Name:       "junitxml",
		ShortHelp:  "Process junit xml files.",
		ShortUsage: "junitxml [COMMAND] [OPTION]... [ARG]...",
		Subcommands: []*cli.Command{
			newMergeCommand(),
			cli.DefaultVersionCommand(os.Stdout),
		},
		Exec: func(ctx context.Context, args []string) error {
			return flag.ErrHelp
		},
	}
}
