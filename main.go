package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func run(filenames []string, op string, col int, stdout io.Writer) error {
	var opFunc statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}
	if col < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, col)
	}
	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)
	// Loop through all files adding their data to consolidate
	for _, fname := range filenames {
		// Open the file for reading
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("Cannot open file: %w", err)
		}
		// Parse the CSV into a slice of float64 numbers
		data, err := cvs2float(f, col)
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}
	_, err := fmt.Fprintln(stdout, opFunc(consolidate))
	return err
}

func main() {
	op := flag.String("op", "", "Operation to be executed. Choices:\nsum - to calculate the sum of all values in the column;\navg - to determine average value of the column.")
	col := flag.Int("col", 0, "csv column on which to execute operation")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s usage information:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	err := run(flag.Args(), *op, *col, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}

}
