package sieve

import (
	"math"
)

func sieveOfEratosthenes(limit int) []int {
	var marks []bool = make([]bool, limit+1)
	var primes []int = make([]int, 0, 10)

	for i := 0; i < len(marks); i++ {
		marks[i] = true
	}

	marks[0] = false
	marks[1] = false

	for i := 2; i*i <= limit; i++ {
		if marks[i] == true {
			for j := i * i; j <= limit; j += i {
				marks[j] = false
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if marks[i] == true {
			primes = append(primes, i)
		}
	}

	return primes
}

func findBiggestSegmentedSieve(left int, right int, primes []int) (int, bool) {
	if left == 1 {
		left = left + 1
	}

	var marks []bool = make([]bool, right-left+1)

	for i := 0; i < len(marks); i++ {
		marks[i] = true
	}

	for i := 0; i < len(primes) && float64(primes[i]) <= math.Sqrt(float64(right)); i++ {
		var base int = (left / primes[i]) * primes[i]
		if base < left {
			base += primes[i]
		}
		for j := base; j <= right; j += primes[i] {
			if j != primes[i] {
				marks[j-left] = false
			}
		}
	}

	for i := right; i >= left; i-- {
		if marks[i-left] == true {
			return i, true
		}
	}

	return -1, false
}

func FindSmallerPrimeNumber(input int) (int, bool) {
	var step int = 100

	right := input - 1

	for {
		left := right - step
		if left < 0 {
			left = 0
		}
		primes := sieveOfEratosthenes(int(math.Sqrt(float64(right))))
		result, ok := findBiggestSegmentedSieve(left, right, primes)

		if ok {
			return result, true
		}

		if !ok && left == 0 {
			break
		}

		right = left
	}

	return -1, false
}
