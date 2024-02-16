package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	data := []float64{34.5, 23.6, 12.56, 78.32}
	exp := 148.98
	res := sum(data)
	if res != exp {
		t.Errorf("Expected %f, got %f instead\n", exp, res)
	}
}

func TestAvg(t *testing.T) {
	data := []float64{34.5, 23.6, 12.56, 78.32}
	exp := 37.245
	res := avg(data)
	if res != exp {
		t.Errorf("Expected %f, got %f instead\n", exp, res)
	}
}

// test all operation functions by applying statsFunc type and table-driven testing
func TestOperations(t *testing.T) {
	data := [][]float64{
		{10, 20, 15, 30, 45, 50, 100, 30},
		{5.5, 8, 2.2, 9.75, 8.45, 3, 2.5, 10.25, 4.75, 6.1, 7.67, 12.287, 5.47},
		{-10, -20},
		{102, 37, 44, 57, 67, 129},
	}

	testCases := []struct {
		name string
		op   statsFunc
		exp  []float64
	}{
		{"Sum", sum, []float64{300, 85.927, -30, 436}},
		{"Avg", avg, []float64{37.5, 6.609769230769231, -15, 72.666666666666666}},
	}

	for _, tc := range testCases {
		for i, d := range data {
			t.Run(tc.name, func(t *testing.T) {
				res := tc.op(d)
				if res != tc.exp[i] {
					t.Errorf("Expected: %v, got: %v", tc.exp[i], res)
				}
			})
		}
	}
}
