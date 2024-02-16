package main

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

// avg determines avarage value of the column
func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}
