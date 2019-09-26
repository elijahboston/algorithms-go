package main

import "testing"

func TestBinarySearch(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	expectedPos := 18
	expectedVal := 67

	actual := BinarySearch(expectedVal, primes)

	if actual != expectedPos {
		t.Errorf("Wrong position got %d expected %d", actual, expectedPos)
	}

	if primes[actual] != expectedVal {
		t.Errorf("Wrong value got %d expected %d", primes[actual], expectedVal)
	}
}
