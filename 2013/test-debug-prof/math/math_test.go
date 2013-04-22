package math

import "testing"

func TestMul(t *testing.T) {
	a, b, expected := 2.0, 3.0, 6.0
	result := Mul(a, b)
	if result != expected {
		t.Fatalf("%v * %v -> %v (expected %v)", a, b, result, expected)
	}
}

func TestDiv(t *testing.T) {
	a, b, expected := 6.0, 3.0, 3.0
	result := Div(a, b)
	if result != expected {
		t.Fatalf("%v / %v -> %v (expected %v)", a, b, result, expected)
	}
}
