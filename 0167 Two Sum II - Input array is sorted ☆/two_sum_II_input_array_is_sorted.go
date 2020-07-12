package main

import "fmt"

func main() {
	numbers, target := []int{2, 7, 11, 15}, 9
	//Output := []int{1, 2}
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	for l, r := 0, len(numbers)-1; l < r; {
		if numbers[l] > target/2 {
			return nil
		}

		if sum := numbers[l] + numbers[r]; sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return nil
}

func twoSumR(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > target/2 {
			return nil
		}

		for j := i + 1; j < len(numbers); j++ {
			if sum := numbers[i] + numbers[j]; sum == target {
				return []int{i + 1, j + 1}
			} else if sum > target {
				break
			} else {
				continue
			}
		}
	}

	return nil
}

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
给定一个已按照升序排列的有序数组，找到两个数使得它们相加之和等于目标数。

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.
函数应该返回这两个下标值index1和index2，其中index1必须小于index2。


Note:
说明：

  Your returned answers (both index1 and index2) are not zero-based.
  返回的下标值（index1和index2）不是从零开始的。

  You may assume that each input would have exactly one solution and you may not use the same element twice.
  你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。


Example:
示例：

  Input: numbers = [2,7,11,15], target = 9
  输入：numbers = [2,7,11,15], target = 9

  Output: [1,2]
  输出：[1,2]

  Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
  解释：2与7之和等于目标数9。因此index1=1，index2=2。
*/
