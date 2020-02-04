package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4, 11, 15}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	resMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := resMap[num]; ok {
			if i > index {
				i, index = index, i
			}

			return []int{i, index}
		}

		key := target - num
		resMap[key] = i
	}

	return nil
}

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

You may assume that each input would have exactly one solution, and you may not use the same element twice.
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。


Example:
示例:

  Given nums = [2, 7, 11, 15], target = 9
  给定 nums = [2, 7, 11, 15], target = 9

  Because nums[0] + nums[1] = 2 + 7 = 9
  因为 nums[0] + nums[1] = 2 + 7 = 9

  return [0, 1]
  所以返回 [0, 1]
*/
