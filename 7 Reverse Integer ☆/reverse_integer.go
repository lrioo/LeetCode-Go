package main

import (
    "fmt"
)

func main() {
    input := 103
    //Output = 301
    fmt.Println(reverse(input))

    input = 120
    //Output = 21
    fmt.Println(reverse(input))

    input = -123
    //Output = -321
    fmt.Println(reverse(input))
}

func reverse(x int) int {
    if x/10 == 0 {
        return x
    }

    var y int
    for ; x/10 != 0; x /= 10 {
        y *= 10
        y += x % 10
    }

    y = y*10 + x
    if y != int(int32(y)) {
        return 0
    }

    return y
}

/*
Given a 32-bit signed integer, reverse digits of an integer.
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

Example 1:
示例 1:

  Input: 123
  输入: 123

  Output: 321
  输出: 321


Example 2:
示例 2:

  Input: -123
  输入: -123

  Output: -321
  输出: -321


Example 3:
示例 3:
  Input: 120
  输入: 120
  Output: 21
  输出: 21


Note:
注意:
    Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the
  purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
    假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
