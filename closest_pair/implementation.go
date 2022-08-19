package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y int
}

func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]point, size)
	for i := 0; i < size; i++ {
		some_point := point{}
		fmt.Printf("Enter x coordinate for %dth point: ", i)
		fmt.Scanf("%d", &some_point.x)
		fmt.Printf("Enter y coordinate for %dth element: ", i)
		fmt.Scanf("%d", &some_point.y)
		arr[i] = some_point
	}
	fmt.Println("Your intial array is: ", arr)

	result := findMinDistance(arr)
	fmt.Println("sorted array is : ", result)
}

func findMinDistance(arr []point) float64 {
	length := len(arr)
	if len(arr) == 0 || len(arr) == 1 {
		return 0
	}
	sorted_by_x := sort_by_x(arr)
	leftMinDistance := findMinDistance(sorted_by_x[:length/2])
	rightMinDistance := findMinDistance(sorted_by_x[length/2:])

	minDistance := math.Min(leftMinDistance, rightMinDistance)
	return combine(minDistance, sorted_by_x)
}

func combine(d float64, arr []point) float64 {
	for true {
		if arr[0].x < (arr[len(arr)/2].x - int(math.Ceil(d))) {
			arr = arr[1:]
		} else {
			break
		}
	}
	for true {
		if arr[len(arr)-1].x < (arr[len(arr)/2].x - int(math.Ceil(d))) {
			arr = arr[:len(arr)-1]
		} else {
			break
		}
	}

	sorted_by_y := sort_by_y(arr)
	return 0
}

func sort_by_x(arr []point) []point {
	length := len(arr)
	if length == 1 {
		return arr
	}
	if length == 2 {
		sort_x(arr)
	}
	return merge(sort_by_x(arr[0:length/2]), sort_by_x(arr[length/2:length]))
}
func sort_by_y(arr []point) []point {
	length := len(arr)
	if length == 1 {
		return arr
	}
	if length == 2 {
		sort_y(arr)
	}
	return merge_y(sort_by_y(arr[0:length/2]), sort_by_y(arr[length/2:length]))
}

func sort_x(arr []point) []point {
	if (len(arr) == 1) || (len(arr) == 2 && (arr[0].x <= arr[1].x)) {
		return arr
	}
	return []point{arr[1], arr[0]}
}
func sort_y(arr []point) []point {
	if (len(arr) == 1) || (len(arr) == 2 && (arr[0].y <= arr[1].y)) {
		return arr
	}
	return []point{arr[1], arr[0]}
}

func merge(arr1, arr2 []point) []point {
	arr3 := make([]point, len(arr1)+len(arr2))
	for i := 0; i < len(arr3); i++ {
		if len(arr2) == 0 {
			copy(arr3[i:], arr1)
			break
		}
		if len(arr1) == 0 {
			copy(arr3[i:], arr2)
			break
		}
		if arr1[0].x < arr2[0].x {
			arr3[i] = arr1[0]
			arr1 = arr1[1:]
			continue
		}
		arr3[i] = arr2[0]
		arr2 = arr2[1:]
	}
	return arr3
}

func merge_y(arr1, arr2 []point) []point {
	arr3 := make([]point, len(arr1)+len(arr2))
	for i := 0; i < len(arr3); i++ {
		if len(arr2) == 0 {
			copy(arr3[i:], arr1)
			break
		}
		if len(arr1) == 0 {
			copy(arr3[i:], arr2)
			break
		}
		if arr1[0].y < arr2[0].y {
			arr3[i] = arr1[0]
			arr1 = arr1[1:]
			continue
		}
		arr3[i] = arr2[0]
		arr2 = arr2[1:]
	}
	return arr3
}
