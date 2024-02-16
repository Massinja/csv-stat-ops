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
