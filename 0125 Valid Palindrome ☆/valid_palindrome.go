package main

import (
	"fmt"
)

func main() {
	Input := "A man, a plan, a canal: Panama"
	//Output := true
	fmt.Println(isPalindrome(Input))

	Input = "race a car"
	//Output := true
	fmt.Println(isPalindrome(Input))
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		hasSpecial := false
		if !isCommonChar(s[i]) {
			i++
			hasSpecial = true
		}
		if !isCommonChar(s[j]) {
			j--
			hasSpecial = true
		}
		if hasSpecial {
			continue
		}

		if s[i] == s[j] || (abs(s[i], s[j]) == 32 && s[i] >= 'A' && s[j] >= 'A') {
			i++
			j--
			continue
		}
		return false
	}

	return true
}

func isCommonChar(c uint8) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func abs(i, j uint8) uint8 {
	if i > j {
		return i-j
	}
	return j-i
}

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

Note: For the purpose of this problem, we define empty string as valid palindrome.
说明：本题中，我们将空字符串定义为有效的回文串。


Example 1:
示例 1:

  Input: "A man, a plan, a canal: Panama"
  输入: "A man, a plan, a canal: Panama"

  Output: true
  输出: true


Example 2:
示例 2:

  Input: "race a car"
  输入: "race a car"

  Output: false
  输出: false


Constraints:
限制条件：

  s consists only of printable ASCII characters.
  s仅包含可打印的ASCII字符。
*/
