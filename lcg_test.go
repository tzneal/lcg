package lcg_test

import (
	"fmt"
	"testing"

	"github.com/tzneal/lcg"
)

func TestNewLCG(t *testing.T) {
	for j := 3; j < 1000; j++ {
		lcg, _ := lcg.NewLCG(j)
		got := map[int]struct{}{}
		for !lcg.Done() {
			got[lcg.Next()] = struct{}{}
		}
		for i := 0; i < j; i++ {
			if _, found := got[i]; !found {
				t.Errorf("expected to find %d", i)
			}
		}
		if len(got) != j {
			t.Errorf("expected %d entries, got %d", j, len(got))
		}
	}
}

func TestLCGSmall(t *testing.T) {
	// test some very small cases that work, but don't necessarily create LCGs
	for j := 1; j <= 2; j++ {
		lcg, _ := lcg.NewLCG(j)
		got := map[int]struct{}{}
		for !lcg.Done() {
			got[lcg.Next()] = struct{}{}
		}
		for i := 0; i < j; i++ {
			if _, found := got[i]; !found {
				t.Errorf("expected to find %d", i)
			}
		}
		if len(got) != j {
			t.Errorf("expected %d entries, got %d", j, len(got))
		}
	}
}

func ExampleLCG() {
	g, _ := lcg.NewLCG(15)
	for !g.Done() {
		fmt.Printf("%d ", g.Next())
	}
	fmt.Println()
	// Output: 8 1 9 2 10 3 11 4 12 5 13 6 14 7 0
}
