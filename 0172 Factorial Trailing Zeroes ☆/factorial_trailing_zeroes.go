package main

import "fmt"

func main() {
	Input := 3
	//Output := 0
	fmt.Println(trailingZeroes(Input))

	Input = 5
	//Output = 1
	fmt.Println(trailingZeroes(Input))
}

func trailingZeroes(n int) int {
	var res int
	for n > 0 {
		n /= 5
		res += n
	}

	return res
}

/*
Given an integer n, return the number of trailing zeroes in n!.
给定一个整数 n，返回 n! 结果尾数中零的数量。


Example 1:
示例1:

  Input: 3
  输入: 3

  Output: 0
  输出: 0

  Explanation: 3! = 6, no trailing zero.
  解释: 3! = 6，尾数中没有零。


Example 2:
示例2:

  Input: 5
  输入: 5

  Output: 1
  输出: 1

  Explanation: 5! = 120, one trailing zero.
  解释: 5! = 120，尾数中有1个零.


Note: Your solution should be in logarithmic time complexity.
说明：你算法的时间复杂度应为O(log n)。
*/
