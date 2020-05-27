package main

import "fmt"

func main() {
	Input := []int{1, 2, 3}
	//Output: [1,2,4]
	fmt.Println(plusOne(Input))

	Input = []int{4, 3, 2, 1}
	//Output: [4,3,2,2]
	fmt.Println(plusOne(Input))

	Input = []int{9, 9, 9}
	//Output: [1,0,0,0]
	fmt.Println(plusOne(Input))

	Input = []int{1, 9, 9}
	//Output: [2,0,0]
	fmt.Println(plusOne(Input))

	Input = []int{0}
	//Output: [1]
	fmt.Println(plusOne(Input))
}

func plusOne(digits []int) []int {
	for i := len(digits); i >= 0; i-- {
		if i == 0 {
			digits = append([]int{1}, digits...)
			return digits
		}

		if digits[i-1]++; digits[i-1] < 10 {
			break
		}
		digits[i-1] = 0
	}
	return digits
}

/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
最高位数字存放在数组的首位，数组中每个元素只存储单个数字。

You may assume the integer does not contain any leading zero, except the number 0 itself.
你可以假设除了整数0之外，这个整数不会以零开头。


Example 1:
示例 1:

  Input:= [1,2,3]
  输入: [1,2,3]

  Output: [1,2,4]
  输出: [1,2,4]

  Explanation: The array represents the integer 123.
  解释: 输入数组表示数字 123。


Example 2:
示例 2:

  Input: [4,3,2,1]
  输入: [4,3,2,1]

  Output: [4,3,2,2]
  输出: [4,3,2,2]

  Explanation: The array represents the integer 4321.
  解释: 输入数组表示数字 4321。
*/
