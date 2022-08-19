// Problem Statement
// Given an array of integers arr, return true if and only if it is a valid mountain array.
// Recall that arr is a mountain array if and only if:
// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
package leetcodearrays

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	tip := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			tip = i
		}
	}
	if tip == 0 {
		return false
	}
	for i := 1; i < tip; i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	for i := tip + 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}
