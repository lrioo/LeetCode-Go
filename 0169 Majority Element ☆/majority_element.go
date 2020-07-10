package main

import "fmt"

func main() {
	Input := []int{3, 2, 3}
	//Output := 3
	fmt.Println(majorityElement(Input))

	Input = []int{2, 2, 1, 1, 1, 2, 2}
	//Output = 2
	fmt.Println(majorityElement(Input))

	Input = []int{2, 2, 2, 1, 1}
	//Output = 2
	fmt.Println(majorityElement(Input))
}

/*
摩尔投票法 Moore Voting —— 需要O(n)的时间和O(1)的空间。
  首先将第一个数字假设为过半数，然后把计数器设为1，比较下一个数和此数是否相等，若相等则计数器加一，反之减一。
  然后看此时计数器的值，若为零，则将下一个值设为候选过半数。
  以此类推直到遍历完整个数组，当前候选过半数即为该数组的过半数。
*/

func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
			cnt++
			continue
		}

		if res == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}

	return res
}

/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
给定一个大小为n的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2⌋ 的元素。

You may assume that the array is non-empty and the majority element always exist in the array.
你可以假设数组是非空的，并且给定的数组总是存在多数元素。


Example 1:
示例1:

  Input: [3,2,3]
  输入: [3,2,3]

  Output: 3
  输出: 3


Example 2:
示例2:

  Input: [2,2,1,1,1,2,2]
  输入: [2,2,1,1,1,2,2]

  Output: 2
  输出: 2
*/
