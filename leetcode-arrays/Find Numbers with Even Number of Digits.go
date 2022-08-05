// Problem Statement
// Given an array nums of integers, return how many of them contain an even number of digits.

package leetcodearrays

func findNumbers(nums []int) int {
	answer := 0
	for _, num := range nums {
		if evenDigits(num) {
			answer++
		}
	}
	return answer
}

func evenDigits(num int) bool {
	if int(num/10) == 0 {
		return false
	}
	if int(num/100) == 0 {
		return true
	}
	if int(num/1000) == 0 {
		return false
	}
	if int(num/10000) == 0 {
		return true
	}
	if int(num/100000) == 0 {
		return false
	}
	return true
}
