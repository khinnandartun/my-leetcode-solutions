package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result []int
	result = append(result, nums1...)
	result = append(result, nums2...)
	sort.Ints(result)
	fmt.Println(result)
	if len(result)%2 == 0 {
		median := float64(result[(len(result)/2)]+result[(len(result)/2)-1]) / 2
		return median
	} else {
		return float64(result[len(result)/2])
	}

}
func main() {
	r := findMedianSortedArrays([]int{1, 3}, []int{2, 4})
	fmt.Println(r)
}
