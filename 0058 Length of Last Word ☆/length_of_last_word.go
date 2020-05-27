package main

import "fmt"

func main() {
	Input := "Hello World"
	//Output := 5
	fmt.Println(lengthOfLastWord(Input))

	Input = "Hello "
	//Output := 5
	fmt.Println(lengthOfLastWord(Input))

	Input = "Hello"
	//Output := 5
	fmt.Println(lengthOfLastWord(Input))

	Input = "H "
	//Output := 1
	fmt.Println(lengthOfLastWord(Input))

	Input = "H"
	//Output := 1
	fmt.Println(lengthOfLastWord(Input))

	Input = "   "
	//Output := 0
	fmt.Println(lengthOfLastWord(Input))

	Input = " "
	//Output := 0
	fmt.Println(lengthOfLastWord(Input))
}

func lengthOfLastWord(s string) int {
	for i := len(s); i > 0; i-- {
		if s[i-1] == ' ' {
			if i != len(s) && i != 1 {
				return len(s) - i
			}

			s = s[:len(s)-1]
		}
	}

	return len(s)
}

func lengthOfLastWord2(s string) int {
	length, f := 0, false
	for i := len(s); i > 0; i-- {
		if s[i-1] == ' ' {
			if f {
				break
			}
			continue
		}

		length++
		f = true
	}

	return length
}

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word
(last word means the last appearing word if we loop from left to right) in the string.
给定一个仅包含大小写字母和空格' '的字符串s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

If the last word does not exist, return 0.
如果不存在最后一个单词，请返回0。

Note: A word is defined as a maximal substring consisting of non-space characters only.
说明：一个单词是指仅由字母组成、不包含任何空格字符的最大子字符串。

Example:
示例:

  Input: "Hello World"
  输入: "Hello World"

  Output: 5
  输出: 5
*/
