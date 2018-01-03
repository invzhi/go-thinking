package prime

import (
	"reflect"
	"testing"
)

func TestGetSlices(t *testing.T) {
	var tests = []struct {
		n      int
		primes []int
	}{
		{5, []int{2, 3, 5}},
		{12, []int{2, 3, 5, 7, 11}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{1, nil},
		{0, nil},
		{-10, nil},
	}

	for _, tt := range tests {
		primes, err := GetSlices(tt.n)
		if err == nil && reflect.DeepEqual(primes, tt.primes) == false {
			t.Errorf("getPrimes(%v) return %v, want %v", tt.n, primes, tt.primes)
		}
	}
}
