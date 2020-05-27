package main

import "fmt"

func main() {
	a, b := "11", "1"
	//Output: "100"
	fmt.Println(addBinary(a, b))

	a, b = "1010", "1011"
	//Output: "10101"
	fmt.Println(addBinary(a, b))

	a, b = "101111", "10"
	//Output: "110001"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	x, carry := []byte(a), uint8('0')
	for i := 1; i < len(x)+1; i++ {
		if i > len(b) {
			if carry == '1' {
				if x[len(x)-i] == '1' {
					x[len(x)-i] = '0'
				} else {
					x[len(x)-i] = '1'
					carry = '0'
				}
				continue
			}

			carry = '0'
			break
		}

		if x[len(x)-i]+b[len(b)-i] == '1'*2 {
			if carry == '1' {
				continue
			} else {
				x[len(x)-i] = '0'
				carry = '1'
			}
		} else if x[len(x)-i]+b[len(b)-i] == '0'*2 {
			x[len(x)-i] = carry
			carry = '0'
		} else {
			if carry == '1' {
				x[len(x)-i] = '0'
			} else {
				x[len(x)-i] = '1'
			}
		}
	}

	if carry > '0' {
		x = append([]byte{carry}, x...)
	}

	return string(x)
}

/*
Given two binary strings, return their sum (also a binary string).
给你两个二进制字符串，返回它们的和（用二进制表示）。

The input strings are both non-empty and contains only characters 1 or 0.
输入为 非空 字符串且只包含数字 1 和 0。


Example 1:
示例 1:

  Input: a = "11", b = "1"
  输入: a = "11", b = "1"

  Output: "100"
  输出: "100"


Example 2:
示例 2:

  Input: a = "1010", b = "1011"
  输入: a = "1010", b = "1011"

  Output: "10101"
  输出: "10101"


Constraints:
提示：

  Each string consists only of '0' or '1' characters.
  每个字符串仅由字符'0'或'1'组成。

  1 <= a.length, b.length <= 10^4
  1 <= a.length, b.length <= 10^4

  Each string is either "0" or doesn't contain any leading zero.
  字符串如果不是"0"，就都不含前导零。
*/
