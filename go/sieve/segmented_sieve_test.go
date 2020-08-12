package sieve

import (
	"testing"
)

func Test_FindSmallerPrimeNumber_SmallNumber(t *testing.T) {
	result, err := FindSmallerPrimeNumber(150)
	if err != nil {
		t.Error("Should not error in finding")
	}

	if result != 149 {
		t.Error("Wrong prime number")
	}
}

func Test_FindSmallerPrimeNumber_ForBigNumber(t *testing.T) {
	result, err := FindSmallerPrimeNumber(49999800)
	if err != nil {
		t.Error("Should not error in finding")
	}

	if result != 49999783 {
		t.Errorf("Wrong prime number: %d", result)
	}
}

func Test_FindSmallerPrimeNumber_ShouldError_WhenInputNumberrEqualOne(t *testing.T) {
	result, err := FindSmallerPrimeNumber(1)
	if err == nil {
		t.Error()
	}

	if result != -1 {
		t.Errorf("Wrong error number: %d", result)
	}
}

func Test_FindSmallerPrimeNumber_ShouldError_WhenInputNumberLowerThanOne(t *testing.T) {
	result, err := FindSmallerPrimeNumber(0)
	if err == nil {
		t.Error()
	}

	if result != -1 {
		t.Errorf("Wrong error number: %d", result)
	}
}