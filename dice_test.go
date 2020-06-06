package main

import (
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
		got := Choose(c.n, c.k)
		if got != c.expected {
			t.Errorf("Choose(%d, %d) == %d, want %d", c.n, c.k, got, c.expected)
		}
	}
}

func TestMultiChoose(t *testing.T) {
	cases := []struct {
		n, k, expected uint64
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 2, 6},
		{4, 2, 10},
		{5, 3, 35},
	}
	for _, c := range cases {
		got := Multichoose(c.n, c.k)
		if got != c.expected {
			t.Errorf("Multichoose(%d, %d) == %d, want %d", c.n, c.k, got, c.expected)
		}
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
		got := CountRollsWithTargetSum(c.n, c.s, c.p)
		if got != c.expected {
			t.Errorf("CountRollsWithTargetSum(%d, %d, %d) == %d, want %d", c.n, c.s, c.p, got, c.expected)
		}
	}
}
