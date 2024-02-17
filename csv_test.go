package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
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

func TestCSV2Float(t *testing.T) {
	csvData := `IP_Address,Requests,Response_Time
192.168.0.199,2056,236
192.168.0.88,899,220
192.168.0.199,3054,226
192.168.0.100,4133,218
192.168.0.199,950,238
`

	testCases := []struct {
		name   string
		col    int
		exp    []float64
		expErr error
		r      io.Reader
	}{
		{name: "Column2",
			col:    2,
			exp:    []float64{2056, 899, 3054, 4133, 950},
			expErr: nil,
			r:      strings.NewReader(csvData),
		},
		{name: "Column3",
			col:    3,
			exp:    []float64{236, 220, 226, 218, 238},
			expErr: nil,
			r:      strings.NewReader(csvData),
		},
		{name: "FailRead",
			col:    1,
			exp:    nil,
			expErr: iotest.ErrTimeout,
			r:      iotest.TimeoutReader(bytes.NewReader([]byte{0})),
		},
		{name: "FailedNotNumber",
			col:    1,
			exp:    nil,
			expErr: ErrNotNumber,
			r:      strings.NewReader(csvData),
		},
		{name: "FailedInvalidColumn1",
			col:    4,
			exp:    nil,
			expErr: ErrInvalidColumn,
			r:      strings.NewReader(csvData),
		},
		{name: "FailedInvalidColumn2",
			col:    -3,
			exp:    nil,
			expErr: ErrInvalidColumn,
			r:      strings.NewReader(csvData),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, resErr := cvs2float(tc.r, tc.col)
			if len(tc.exp) == 0 && len(res) != 0 {
				t.Errorf("Expected an empty slice; got: %v", res)
				return
			} else if len(res) == 0 && len(tc.exp) != 0 {
				t.Errorf("Received an empty slice; expected: %v", tc.exp)
				return
			} else if len(tc.exp) != len(res) {
				t.Errorf("Slice length is wrong. Expected %d elements; got: %v", len(tc.exp), len(res))
				return
			}
			for i, exp := range tc.exp {
				if res[i] != exp {
					t.Errorf("Expected: %v; got: %v", exp, res[i])
				}
			}
			if !errors.Is(resErr, tc.expErr) {
				t.Errorf("Expected error: %v; got: %v", tc.expErr, resErr)
			}
		})
	}
}
