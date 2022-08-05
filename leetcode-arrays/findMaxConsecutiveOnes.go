// Problem Statement
// Given a binary array nums, return the maximum number of consecutive 1's in the array.

package leetcodearrays

func findMaxConsecutiveOnes(nums []int) int {
	answer := 0
	last_zero := -1
	now_zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			now_zero = i
			if now_zero-last_zero > answer {
				answer = now_zero - last_zero - 1
			}
			last_zero = now_zero
		}
		if i == (len(nums)-1) && nums[i] == 1 {
			now_zero = i
			if now_zero-last_zero > answer {
				answer = now_zero - last_zero
			}
			last_zero = now_zero
		}
	}
	return answer
}
