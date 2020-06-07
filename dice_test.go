package main

import (
	"math"
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		n        uint64
		expected *big.Int
	}{
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(6)},
		{4, big.NewInt(24)},
		{8, big.NewInt(40320)},
	}
	for _, c := range cases {
		got := Factorial(c.n)
		if got.Cmp(c.expected) != 0 {
			t.Errorf("Factorial(%d) == %d, want %d", c.n, got, c.expected)
		}
	}
}

func TestChoose(t *testing.T) {
	cases := []struct {
		n, k, expected uint64
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 2, 3},
		{4, 2, 6},
		{4, 4, 1},
		{9, 5, 126},
		{10, 2, 45},
		{10, 3, 120},
		{12, 3, 220},
		{18, 9, 48620},
		{30, 21, 14307150},
	}
	for _, c := range cases {
		got, _ := Choose(c.n, c.k)
		if got != c.expected {
			t.Errorf("Choose(%d, %d) == %d, want %d", c.n, c.k, got, c.expected)
		}
	}
}

func TestPermutations(t *testing.T) {
	cases := []struct {
		s, n, expected uint64
	}{
		{6, 1, 6},
		{6, 2, 36},
		{10, 3, 1000},
	}
	for _, c := range cases {
		got, _ := Permutations(c.s, c.n)
		if got != c.expected {
			t.Errorf("Permutations(%d, %d) == %d, want %d", c.s, c.n, got, c.expected)
		}

	}

	// testing error case
	_, err := Permutations(2, 65)
	if err == nil {
		t.Error("Permutations(2, 65) did not result in error")
	}
}

func TestCountRollsWithTargetSum(t *testing.T) {
	cases := []struct {
		p, n, s, expected uint64
	}{
		{5, 1, 6, 1},
		{2, 2, 6, 1},
		{3, 2, 6, 2},
		{7, 2, 6, 6},
		{31, 10, 6, 3393610},
	}
	for _, c := range cases {
		got, _ := CountRollsWithTargetSum(c.n, c.s, c.p)
		if got != c.expected {
			t.Errorf("CountRollsWithTargetSum(%d, %d, %d) == %d, want %d", c.n, c.s, c.p, got, c.expected)
		}
	}
}

func TestCountRollsGreaterOrEqualToTargetSum(t *testing.T) {
	cases := []struct {
		p, n, s, expected uint64
	}{
		{12, 2, 6, 1},
		{11, 2, 6, 3},
		{10, 2, 6, 6},
	}
	for _, c := range cases {
		got, _ := CountRollsGreaterOrEqualToTargetSum(c.n, c.s, c.p)
		if got != c.expected {
			t.Errorf("CountRollsGreaterOrEqualToTargetSum(%d, %d, %d) == %d, want %d", c.n, c.s, c.p, got, c.expected)
		}
	}
}

func TestChanceToMatchOrBeat(t *testing.T) {
	cases := []struct {
		num, sides, target uint64
		expected           float64
	}{
		{2, 6, 5, 83.33},
		{3, 6, 10, 62.50},
		{3, 6, 15, 9.26},
		{4, 6, 15, 44.37},
		{4, 5, 20, 5.4},
	}
	for _, c := range cases {
		got, _ := ChanceToMatchOrBeat(c.num, c.sides, c.target)
		if AlmostEq(got, c.expected) {
			t.Errorf("ChanceToMatchOrBeat(%d, %d, %d) == %f, want %f", c.num, c.sides, c.target, got, c.expected)
		}
	}
}

// utility function
func AlmostEq(left float64, right float64) bool {
	tolerance := 0.001
	diff := math.Abs(left - right)
	return diff <= tolerance
}
