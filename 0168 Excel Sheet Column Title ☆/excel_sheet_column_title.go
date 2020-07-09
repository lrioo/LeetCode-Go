package main

import (
	"fmt"
)

func main() {
	Input := 1
	//Output := "A"
	fmt.Println(convertToTitle(Input))

	Input = 28
	//Output = "AB"
	fmt.Println(convertToTitle(Input))

	Input = 701
	//Output = "ZY"
	fmt.Println(convertToTitle(Input))

	Input = 703
	//Output = "ZY"
	fmt.Println(convertToTitle(Input))
}

func convertToTitle(n int) string {
	var res string
	for n > 0 {
		n--
		res = string(n%26+'A') + res
		n /= 26
	}

	return res
}

/*
Given a positive integer, return its corresponding column title as appear in an Excel sheet.
给定一个正整数，返回它在Excel表中相对应的列名称。

For example:
例如：

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...


Example 1:
示例1:

  Input: 1
  输入: 1

  Output: "A"
  输出: "A"


Example 2:
示例2:

  Input: 28
  输入: 28

  Output: "AB"
  输出: "AB"


Example 3:
示例3:

  Input: 701
  输入: 701

  Output: "ZY"
  输出: "ZY"
*/
