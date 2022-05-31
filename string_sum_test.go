package string_sum

import (
	"testing"
)

type testPairSlice struct {
	str string
	sum string
}

func TestStringSum(t *testing.T) {
	var test testPairSlice = testPairSlice{"+2 - -3", "5"}

	res, err := StringSum(test.str)

	if err != nil {
		t.Error(err)
	}

	if res != test.sum {
		t.Error("Expected ", test.sum,
			"got ", res)
	}
}
