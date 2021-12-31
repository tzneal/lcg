package lcg_test

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/tzneal/lcg"
)

func TestIsPrime(t *testing.T) {
	for i := 2; i < 1e4; i++ {
		bi := big.NewInt(int64(i))
		truth := bi.ProbablyPrime(10)
		if got := lcg.IsPrime(i); got != truth {
			t.Errorf("IsPrime(%d) = %v, want %v", i, got, truth)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2}},
		{8, []int{2}},
		{6, []int{2, 3}},
		{15, []int{3, 5}},
		{42, []int{2, 3, 7}},
		{49, []int{7}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := lcg.PrimeFactors(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
