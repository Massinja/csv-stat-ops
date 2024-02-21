package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
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

	// Create the channels to receive results or errors of operations
	resCh := make(chan []float64)
	errCh := make(chan error)

	// doneCh won't send any data, only a signal indicating the process is done
	doneCh := make(chan struct{})

	// coordinate goroutines execution
	wg := sync.WaitGroup{}

	// Loop through all files and create a goroutine to process
	for _, fname := range filenames {
		wg.Add(1)
		go func(fname string) {

			defer wg.Done()

			// Open the file for reading
			f, err := os.Open(fname)
			if err != nil {
				errCh <- fmt.Errorf("Cannot open file: %w", err)
				return
			}
			// Parse the CSV into a slice of float64 numbers
			data, err := cvs2float(f, col)
			if err != nil {
				errCh <- err
			}
			if err := f.Close(); err != nil {
				errCh <- err
			}
			// Append the data to consolidate
			resCh <- data
		}(fname)
	}
	go func() {
		wg.Wait()
		close(doneCh)
	}()
	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidate = append(consolidate, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(stdout, opFunc(consolidate))
			return err
		}
	}
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
