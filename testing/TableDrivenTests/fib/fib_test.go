package fib_test

import (
	"TestDrivenTests/fib"
	"testing"
)

func TestFibonacсi(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"first", 1, 1},
		{"second", 2, 1},
		{"third", 3, 2},
		{"fourth", 4, 3},
	}

	for _, tc := range testCases {
		res := fib.Fibonacсi(tc.n)
		check := res == tc.want
		if !check {
			t.Errorf("test %s failed: got: %d, want: %d", tc.name, res, tc.want)
		}
	}
}
func ExampleFibonacсi() {
	x := 4
	fib.Fibonacсi(x)
	y := 3
	fib.Fibonacсi(y)
	//Output
	//3
	//2
}
