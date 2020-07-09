package main

import (
	"fmt"
)

func main() {
	Input := "A"
	//Output := 1
	fmt.Println(titleToNumber(Input))

	Input = "AB"
	//Output = 28
	fmt.Println(titleToNumber(Input))

	Input = "ZY"
	//Output = 701
	fmt.Println(titleToNumber(Input))
}

func titleToNumber(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*26 + int(s[i]-'A'+1)
	}
	return res
}

/*
Given a column title as appear in an Excel sheet, return its corresponding column number.
给定一个Excel表格中的列名称，返回其相应的列序号。

For example:
例如：

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...


Example 1:
示例 1:

  Input: "A"
  输入: "A"

  Output: 1
  输出: 1


Example 2:
示例 2:

  Input: "AB"
  输入: "AB"

  Output: 28
  输出: 28


Example 3:
示例 3:

  Input: "ZY"
  输入: "ZY"

  Output: 701
  输出: 701


Constraints:
提示：

  1. 1 <= s.length <= 7
  1. 1 <= s.length <= 7。

  2. s consists only of uppercase English letters.
  2. s仅包含大写英文字母。

  3. s is between "A" and "FXSHRXW".
  3. s在"A"和"FXSHRXW"之间。
*/
