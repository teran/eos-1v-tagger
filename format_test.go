package tagger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatter(t *testing.T) {
	r := require.New(t)

	type testCase struct {
		name      string
		sample    string
		subst     map[string]interface{}
		expResult string
	}

	tcs := []testCase{
		{
			name:      "float64",
			sample:    `TEST_${testvar:.2f}_blah`,
			subst:     map[string]interface{}{"testvar": 1.24},
			expResult: `TEST_1.24_blah`,
		},
		{
			name:      "float64 with 1 number precision",
			sample:    `TEST_${testvar:.1f}_blah`,
			subst:     map[string]interface{}{"testvar": 1.2455433},
			expResult: `TEST_1.2_blah`,
		},
		{
			name:      "int64 with leading zeros",
			sample:    `TEST_${testvar:05d}_blah`,
			subst:     map[string]interface{}{"testvar": 2},
			expResult: `TEST_00002_blah`,
		},
		{
			name:      "string",
			sample:    `TEST${testvar:s}blah`,
			subst:     map[string]interface{}{"testvar": "blah"},
			expResult: `TESTblahblah`,
		},
		{
			name:      "multiple variables",
			sample:    `TEST_${var1:s}_${var2:d}_test`,
			subst:     map[string]interface{}{"var1": "test1", "var2": 3975},
			expResult: `TEST_test1_3975_test`,
		},
	}

	for _, tc := range tcs {
		res := Format(tc.sample, tc.subst)
		r.Equalf(tc.expResult, res, tc.name)
	}
}
