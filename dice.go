package main

import (
	"errors"
	"fmt"
	"math/big"
)

// Combinatorics, which are featured prominently in probablility in dice
// quickly results in very large numbers, due to the presence of
// factorials in this field of math.
// For this reason we are performing many calculations using math/big

// Common value used in comparisons
var zero = big.NewInt(0)
var one = big.NewInt(1)

// Dynamic Programming: Calculated values are saved to reduce duplicate effort
var facts [41]big.Int

func main() {

	fmt.Println("What are the odds?")
}

func CountRollsGreaterOrEqualToTargetSum(n uint64, s uint64, p uint64) (uint64, error) {
	var total uint64
	// fmt.Printf("from %d to %d\n", p, s*n)
	for i := p; i <= s*n; i++ {
		// fmt.Printf("CountRollsWithTargetSum(%d, %d, %d)\n", n, s, i)
		addend, err := CountRollsWithTargetSum(n, s, i)
		if err != nil {
			return 0, err
		}
		total += addend
	}
	return total, nil
}

// Probability P of getting a sum p by rolling n dice each with s sides
//   is expressed by the formula
//   P(p,n,s) = (1/math.Pow(s, n)) * sum[k: 0, math.Floor((p-n)/s)](math.Pow(-1, k) * Choose(n, k) * Choose((p - s * k - 1), (p - s * k - n) ) )
//     p = "points" or the target sum of the result of all dice rolled
//     n = number of dice to roll
//     s = sides on each of n dice (all dice assumed to have the same number of sides)
//     with credit to https://www.lucamoroni.it/the-dice-roll-sum-problem/
// This calculation is broken down into subcalculations by functions defined below to answer the questions proposed

// Count number of possible outcomes from rolling n dice, each with s sides
func Permutations(s uint64, n uint64) (uint64, error) {
	sides := big.NewInt(0).SetUint64(s)
	num := big.NewInt(0).SetUint64(n)
	permutations := big.NewInt(0).Exp(sides, num, zero)
	if permutations.IsUint64() {
		return permutations.Uint64(), nil
	}
	return 0, errors.New(fmt.Sprintf("Value %d**%d is too large for uint64", s, n))
}

// Count the number of ways to get target sum from a given number of dice, each with given sides
//   expressed by sum[k: 0, math.Floor((p-n)/s)](math.Pow(-1, k) * Choose(n, k) * Choose((p - s * k - 1), (p - s * k - n) ))
//   p becomes target
func CountRollsWithTargetSum(n uint64, s uint64, p uint64) (uint64, error) {
	var total uint64
	k := (p - n) / s
	for i := uint64(0); i <= k; i++ {
		first, err := Choose(n, i)
		if err != nil {
			return 0, err
		}
		second, err := Choose(p-s*i-1, p-s*i-n)
		if err != nil {
			return 0, err
		}

		// we just add or subtract below since math.Pow(-1, k) is always 1 or -1
		if shouldAdd(i) {
			total += first * second
		} else {
			total -= first * second
		}
	}
	return total, nil
}

// Combinations without Repetitions
//   commonly expressed "n choose k"
func Choose(n uint64, k uint64) (uint64, error) {
	if k == 1 {
		return n, nil
	}
	if k == 0 || k == n {
		return 1, nil
	}

	numerator := Factorial(n)
	denominator := big.NewInt(0).Mul(Factorial(k), Factorial(n-k))

	var result big.Int
	result.Div(numerator, denominator)
	if result.IsUint64() {
		return result.Uint64(), nil
	}
	return 0, errors.New(fmt.Sprintf("Value \"%d choose %d\" is too large for uint64", n, k))
}

// Calculate the factorial for a value
//   uses the global facts array to recall previously-calculated values
//   the value of approximately 21! and beyond exceed uint64 so math.big.Int is returned
func Factorial(num uint64) *big.Int {
	facts[1] = *one

	if facts[num].Cmp(zero) == 0 {
		for i := uint64(2); i <= num; i++ {
			if facts[i].Cmp(zero) == 0 {
				facts[i].Mul(&facts[i-1], big.NewInt(0).SetUint64(i))
			}
		}
	}
	return big.NewInt(0).Set(&facts[num])

}

// The first term of the sum of series in CountRollsWithTargetSum is always 1 or -1, based on
//   whether the current term index is even or not. Here we accept that index and hint to
//   CountRollsWithTargetSum whether to add or subtract the next term from the running total.
func shouldAdd(in uint64) bool {
	if in%2 == 0 {
		return true
	}
	return false
}
