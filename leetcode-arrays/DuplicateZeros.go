// Problem Statement
// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
// Note that elements beyond the length of the original array are not written.
// Do the above modifications to the input array in place and do not return anything.

package leetcodearrays

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			//arr = append(arr,0)
			copy(arr[i+2:], arr[i+1:])
			arr[i+1] = 0
			i++
		}
	}
}
