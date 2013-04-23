package math

import "testing"

func checkMul(t *testing.T, a, b, expected float64) {
	t.Logf("Testing Mul(%f, %f)", a, b)
	if Mul(a, b) != expected {
		t.Fatalf("Mul(%f, %f) failed", a, b) // FIXME: Better error
	}
}

var mulTestCases = []struct {
	a, b, expected float64
}{
	{1., 2., 2.},
	{-2., -3., 6.},
}

func TestMul(t *testing.T) {
	for _, tc := range mulTestCases {
		checkMul(t, tc.a, tc.b, tc.expected)
	}
}
