package math

import "testing"

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ { // go test will increase N to get good timing
		Div(10.0, 3.0)
	}
}
