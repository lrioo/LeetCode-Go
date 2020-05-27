package main

import "fmt"

func main() {
	Input := 4
	//Output: 2
	fmt.Println(mySqrt(Input))

	Input = 8
	//Output: 2
	fmt.Println(mySqrt(Input))

	Input = 9
	//Output: 3
	fmt.Println(mySqrt(Input))

	Input = 0
	//Output: 0
	fmt.Println(mySqrt(Input))

	Input = 15
	//Output: 3
	fmt.Println(mySqrt(Input))
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 0, x
	for l < r {
		m := l + (r-l)/2
		if x/m >= m {
			l = m + 1
		} else {
			r = m
		}
	}

	return r - 1
}

func mySqrt2(x int) int {
	if x < 2 {
		return x
	}

	l, r := 0, x
	for l < r {
		m := l + (r-l)/2
		if x/m >= m {
			l = m + 1
		} else {
			r = m
		}
	}

	return r - 1
}

/*
Implement int sqrt(int x).
实现 int sqrt(int x) 函数。

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
计算并返回 x 的平方根，其中 x 是非负整数。

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。


Example 1:
示例 1:

  Input: 4
  输入: 4

  Output: 2
  输出: 2


Example 2:
示例 2:

  Input: 8
  输入: 8

  Output: 2
  输出: 2


Explanation:
说明:

  The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.
  8的平方根是2.82842...，由于返回类型是整数，小数部分将被舍去。

*/
