/*
First Position of Target

Description
Given a sorted nums (ascending order) and a target number, find the first index of this number in O(log n) time complexity.

If the target number does not exist in the nums, return -1.
*/

package main

import "fmt"

const (
	NOFOUND       = -1
	INVALID_PARAM = -1
)

var (
	// nums = []int{1, 2, 3, 4, 5} // ascending
	// nums = []int{}
	nums = []int{2, 2, 3, 4, 5}
	// nums = []int{1, 3, 3, 4, 5}
	target = 0
)

func FPT(nums []int, target int) int {
	if len(nums) == 0 {
		return INVALID_PARAM
	}
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid
		}
		if nums[mid] >= target {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return NOFOUND
}

func main() {
	fmt.Print(FPT(nums, target))
}
