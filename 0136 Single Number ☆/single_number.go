package main

import "fmt"

func main() {
	Input := []int{2, 2, 1}
	//Output := 1
	fmt.Println(singleNumber(Input))

	Input = []int{4, 1, 2, 1, 2}
	//Output := 4
	fmt.Println(singleNumber(Input))
}

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}

	return nums[0]
}

/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

Note:
说明：

  Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
  你的算法应该具有线性时间复杂度。你可以不使用额外空间来实现吗？


Example 1:
示例1：

  Input: [2,2,1]
  输入：[2,2,1]

  Output：1
  输出：1


Example 2:
示例2：

  Input: [4,1,2,1,2]
  输入：[4,1,2,1,2]

  Output：4
  输出：4
*/
