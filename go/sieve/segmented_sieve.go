package sieve

import (
	"errors"
	"math"
)

func findPrimesBySieveOfEratosthenes(limit int) []int {
	var isPrimeMarks []bool = make([]bool, limit+1)
	var primes []int = make([]int, 0, 10)

	for i := 0; i < len(isPrimeMarks); i++ {
		isPrimeMarks[i] = true
	}

	isPrimeMarks[0] = false
	isPrimeMarks[1] = false

	for i := 2; i*i <= limit; i++ {
		if isPrimeMarks[i] == true {
			for j := i * i; j <= limit; j += i {
				isPrimeMarks[j] = false
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if isPrimeMarks[i] == true {
			primes = append(primes, i)
		}
	}

	return primes
}

func findBiggestPrimeInSegmentedSieve(left int, right int, initialPrimes []int) (int, bool) {
	if left == 1 {
		left = left + 1
	}

	var isPrimeMarks []bool = make([]bool, right-left+1)

	for i := 0; i < len(isPrimeMarks); i++ {
		isPrimeMarks[i] = true
	}

	for i := 0; i < len(initialPrimes) && float64(initialPrimes[i]) <= math.Sqrt(float64(right)); i++ {
		var procesingPrime = initialPrimes[i]
		var firstMultipleOfProcessingPrime int = (left / procesingPrime) * procesingPrime

		if firstMultipleOfProcessingPrime < left {
			firstMultipleOfProcessingPrime += procesingPrime
		}

		for j := firstMultipleOfProcessingPrime; j <= right; j += procesingPrime {
			if j != procesingPrime {
				isPrimeMarks[j-left] = false
			}
		}
	}

	for i := right; i >= left; i-- {
		if isPrimeMarks[i-left] == true {
			return i, true
		}
	}

	return -1, false
}

// FindSmallerPrimeNumber help to find first smaller number of the input
func FindSmallerPrimeNumber(input int) (int, error) {
	if input <= 1 {
		return -1, errors.New("Input should bigger than 1")
	}

	var step int = 100

	right := input - 1

	for {
		left := right - step
		if left < 0 {
			left = 0
		}
		primes := findPrimesBySieveOfEratosthenes(int(math.Sqrt(float64(right))))
		result, ok := findBiggestPrimeInSegmentedSieve(left, right, primes)

		if ok {
			return result, nil
		}

		if !ok && left == 0 {
			break
		}

		right = left
	}

	return -1, errors.New("Unable to find lower prime number")
}
