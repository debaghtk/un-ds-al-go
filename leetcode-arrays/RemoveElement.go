// Problem Statement
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.

package leetcodearrays

func removeElement(nums []int, val int) int {
	answer := len(nums)
	for i := 0; i < answer; i = i {
		if nums[i] == val {
			answer--
			copy(nums[i:], nums[i+1:])
			continue
		}
		i++
	}
	return answer
}
