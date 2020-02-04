package main

import (
	"fmt"
)

func main() {
	Input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//Output := 6
	fmt.Println(maxSubArray(Input))
}

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for _, num := range nums[1:] {
		if sum += num; sum < num {
			sum = num
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。


Example:
示例:

  Input: [-2,1,-3,4,-1,2,1,-5,4],
  输入: [-2,1,-3,4,-1,2,1,-5,4],

  Output: 6
  输出: 6

  Explanation: [4,-1,2,1] has the largest sum = 6.
  解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


Follow up:
进阶:

  If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
  如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
