package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var svc = NewPacksService()

type testCase struct {
	name           string      // test case name
	input          int         // function input
	expectedResult map[int]int // expected outcome
	hasError       bool        // a flag to check that an error is expected
}

func TestCalculateExactAmount(t *testing.T) {
	tc := testCase{
		name:  "test exact amount",
		input: 300,
		expectedResult: map[int]int{
			100: 3,
		},
		hasError: false,
	}

	actualResult, _ := svc.Calculate(tc.input)
	assert.Equal(t, tc.expectedResult, actualResult)
}

func TestCalculateInBetween(t *testing.T) {
	tc := testCase{
		name:  "test in between two pack sizes added, but less than second pack x2",
		input: 399,
		expectedResult: map[int]int{
			100: 4,
		},
		hasError: false,
	}

	actualResult, _ := svc.Calculate(tc.input)
	assert.Equal(t, tc.expectedResult, actualResult)
}

func TestCalculateBigNumber(t *testing.T) {
	tc := testCase{
		name:  "test in between two pack sizes added, but less than second pack x2",
		input: 1000,
		expectedResult: map[int]int{
			500: 2,
		},
		hasError: false,
	}

	actualResult, _ := svc.Calculate(tc.input)
	assert.Equal(t, tc.expectedResult, actualResult)
}

func TestCalculateBigNumberAndPipi(t *testing.T) {
	tc := testCase{
		name:  "test in between two pack sizes added, but less than second pack x2",
		input: 1001,
		expectedResult: map[int]int{
			500: 1,
			100: 3,
			250: 1,
		},
		hasError: false,
	}

	actualResult, _ := svc.Calculate(tc.input)
	assert.Equal(t, tc.expectedResult, actualResult)
}
