package main

import (
    "fmt"
)

func main() {
    input := 121
    //Output = true
    fmt.Println(isPalindrome(input))

    input = -121
    //Output = false
    fmt.Println(isPalindrome(input))

    input = 10
    //Output = false
    fmt.Println(isPalindrome(input))
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    if x < 10 {
        return true
    }

    if x%10 == 0 {
        return false
    }

    var y int
    for y < x {
        y = y*10 + x%10
        x /= 10
        if y == x || y == x/10 {
            return true
        }
    }
    return false
}

/* 解法二: 转换为字符串进行比较
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    s := []rune(strconv.FormatInt(int64(x), 10))
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }

    return true
}
*/

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

Example 1:
示例 1:

  Input: 121
  输入: 121

  Output: true
  输出: true


Example 2:
示例 2:

  Input: -121
  输入: -121

  Output: false
  输出: false

  Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
  解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。


Example 3:
示例 3:

  Input: 10
  输入: 10

  Output: false
  输出: false

  Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
  解释: 从右向左读, 为 01 。因此它不是一个回文数。


Follow up:
进阶:

  Could you solve it without converting the integer to a string?
  你能不将整数转为字符串来解决这个问题吗？


Hint:
提示:

  Beware of overflow when you reverse the integer.
  反转整数时要小心溢出。
*/
