package main

import "fmt"

func main() {
	Input := 2
	//Output: 2
	fmt.Println(climbStairs(Input))

	Input = 3
	//Output: 3
	fmt.Println(climbStairs(Input))

	Input = 1
	//Output: 1
	fmt.Println(climbStairs(Input))

	Input = 5
	//Output: 5
	fmt.Println(climbStairs(Input))
}

func climbStairs(n int) int {
	a, b := 1, 1
	for ; n > 0; n-- {
		b += a
		a = b - a
	}

	return a
}

/*
You are climbing a stair case. It takes n steps to reach to the top.
假设你正在爬楼梯。需要n阶你才能到达楼顶。

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
每次你可以爬1或2个台阶。你有多少种不同的方法可以爬到楼顶呢？

Note: Given n will be a positive integer.
注意：给定n是一个正整数。

Example 1:
示例1：

  Input: 2
  输入：2

  Output: 2
  输出：2

Explanation: There are two ways to climb to the top.
解释：有两种方法可以爬到楼顶。

  1. 1 step + 1 step
  1. 1 阶 + 1 阶

  2. 2 steps
  2. 2阶


Example 2:
示例2：

  Input: 3
  输入：3

  Output: 3
  输出：3

Explanation: There are three ways to climb to the top.
解释：有三种方法可以爬到楼顶。

  1. 1 step + 1 step + 1 step
  1. 1阶 + 1阶 + 1阶

  2. 1 step + 2 steps
  2. 1阶 + 2阶

  3. 2 steps + 1 step
  3. 2阶 + 1阶
*/
