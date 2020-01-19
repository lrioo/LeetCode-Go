package main

import "fmt"

func main() {
	//Input: [1, 3, 5, 7], 0
	//Output: 0
	nums, target := []int{1, 3, 5, 7}, 0
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 1
	//Output: 0
	nums, target = []int{1, 3, 5, 7}, 1
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 2
	//Output: 1
	nums, target = []int{1, 3, 5, 7}, 2
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 3
	//Output: 1
	nums, target = []int{1, 3, 5, 7}, 3
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 4
	//Output: 2
	nums, target = []int{1, 3, 5, 7}, 4
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 5
	//Output: 2
	nums, target = []int{1, 3, 5, 7}, 5
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 6
	//Output: 3
	nums, target = []int{1, 3, 5, 7}, 6
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 7], 7
	//Output: 3
	nums, target = []int{1, 3, 5, 7}, 7
	fmt.Println(searchInsert(nums, target))

	//Input: [1, 3, 5, 6], 8
	//Output: 4
	nums, target = []int{1, 3, 5, 6}, 8
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	pre, end := 0, len(nums)-1
	for pre < end {
		idx := (pre + end) / 2

		if target == nums[idx] {
			return idx
		}

		if target > nums[idx] {
			pre = idx + 1
		}

		if target < nums[idx] {
			end = idx
		}
	}

	return pre
}

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

You may assume no duplicates in the array.
你可以假设数组中无重复元素。


Example 1:
示例 1:

  Input: [1,3,5,6], 5
  输入: [1,3,5,6], 5

  Output: 2
  输出: 2


Example 2:
示例 2:

  Input: [1,3,5,6], 2
  输入: [1,3,5,6], 2

  Output: 1
  输出: 1


Example 3:
示例 3:

  Input: [1,3,5,6], 7
  输入: [1,3,5,6], 7

  Output: 4
  输出: 4


Example 4:
示例 4:

  Input: [1,3,5,6], 0
  输入: [1,3,5,6], 0

  Output: 0
  输出: 0
*/
