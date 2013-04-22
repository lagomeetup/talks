package table

import "testing"

func checkDouble(t *testing.T, value, expected int) {
	t.Logf("Testing Double(%d)\n", value)
	result := Double(value)
	if result != expected {
		t.Fatalf("Double(%d) = %d! (should have been %d)", value, result, expected)
	}
}

type doubleTestCase struct {
	value, expected int
}

var doubleTestCases = []doubleTestCase {
	{1, 2},
	{-1, -2},
}

func TestDouble(t *testing.T) {
	for _, tc := range doubleTestCases {
		checkDouble(t, tc.value, tc.expected)
	}
}
