package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {

	testCases := []struct {
		name   string
		files  []string
		op     string
		col    int
		expErr error
		exp    string
	}{
		{name: "SumSecondColumn",
			files:  []string{"./testdata/data.csv"},
			op:     "sum",
			col:    2,
			expErr: nil,
			exp:    "11092\n",
		},
		{name: "SumTwoFiles",
			files:  []string{"./testdata/data.csv", "./testdata/data2.csv"},
			op:     "sum",
			col:    2,
			expErr: nil,
			exp:    "38356\n",
		},
		{name: "NoFiles",
			files:  []string{},
			op:     "avg",
			col:    2,
			expErr: ErrNoFiles,
			exp:    "",
		},
		{name: "AvgNoColumn",
			files:  []string{"./testdata/data.csv"},
			op:     "avg",
			col:    0,
			expErr: ErrInvalidColumn,
			exp:    "",
		},
		{name: "AvgNoOperation",
			files:  []string{"./testdata/data.csv"},
			op:     "",
			col:    2,
			expErr: ErrInvalidOperation,
			exp:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &res)
			if tc.expErr == nil && err != nil {
				t.Errorf("Expected no errors; got: %v", err)
			}
			if !errors.Is(err, tc.expErr) {
				t.Errorf("Expected error: %v; got: %v", tc.expErr, err)
			}
			if res.String() != tc.exp {
				t.Errorf("Expected: %v; got: %v", tc.exp, &res)
			}
		})
	}
}
