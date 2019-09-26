package main

import "math"

func search(n int, arr []int, min int, max int) int {
	if min == max {
		return -1
	}

	estimated := float64((min + max) / 2)
	pos := int(math.Round(estimated))
	val := arr[pos]

	if val == n {
		return pos
	}

	if val > n {
		return search(n, arr, min, pos)
	}

	if val < n {
		return search(n, arr, pos, max)
	}

	return -1
}

// BinarySearch performs a search on a given array for "n"
func BinarySearch(n int, arr []int) int {
	return search(n, arr, 0, len(arr))
}
