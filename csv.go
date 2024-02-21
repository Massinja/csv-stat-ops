package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

// sum calculates the sum of all values in the column
func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

// avg determines average value of the column
func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// cvs2float parses the contents of the csv file into a []float64
// parses the contents of the specified column only
func cvs2float(r io.Reader, column int) ([]float64, error) {
	if column < 1 {
		return nil, fmt.Errorf("%w", ErrInvalidColumn)
	}

	// adjusting for a 0 based index
	column--

	cr := csv.NewReader(r)
	// ReuseRecord allows reader to reuse the same slice for each read operation
	cr.ReuseRecord = true

	var data []float64

	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot read data from a file: %w", err)
		}
		// discard  the title line
		if i == 0 {
			continue
		}
		if column >= len(row) {
			return nil, fmt.Errorf("%w: file has only %d columns", ErrInvalidColumn, len(row))
		}
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	return data, nil
}
