package main

import (
	"context"
	"encoding/xml"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/cluttrdev/cli"

	"go.cluttr.dev/junitxml"
)

func newMergeCommand() *cli.Command {
	fs := flag.NewFlagSet("junitxml merge", flag.ExitOnError)

	output := fs.String("o", "", "Write to file instead of stdout")

	return &cli.Command{
		Name:       "merge",
		ShortHelp:  "Merge junit xml files",
		ShortUsage: "junitxml merge [-o <file>] [FILE]...",
		Flags:      fs,
		Exec: func(ctx context.Context, args []string) error {
			if len(args) == 0 {
				return flag.ErrHelp
			}

			filepaths := args

			reports := make([]junitxml.TestReport, 0, len(filepaths))
			for _, filepath := range filepaths {
				file, err := os.Open(filepath)
				if err != nil {
					slog.Error("failed to open file", "error", err, "file", filepath)
					continue
				}
				defer func() {
					_ = file.Close()
				}()

				report, err := junitxml.Parse(file)
				if err != nil {
					slog.Error("failed to parse file", "error", err, "file", filepath)
					continue
				}

				reports = append(reports, report)
			}

			report := junitxml.Merge(reports)

			var ofile *os.File
			if *output != "" {
				var err error
				ofile, err = os.Create(*output)
				if err != nil {
					return fmt.Errorf("create output file: %w", err)
				}
			} else {
				ofile = os.Stdout
			}

			encoder := xml.NewEncoder(ofile)
			encoder.Indent("", "\t")
			if err := encoder.Encode(report); err != nil {
				return fmt.Errorf("write file: %w", err)
			}

			return nil
		},
	}
}
