package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Input = 1
	//Output = 1
	Input := 1
	fmt.Println(countAndSay(Input))

	//Input = 2
	//Output = 11
	Input = 2
	fmt.Println(countAndSay(Input))

	//Input = 3
	//Output = 21
	Input = 3
	fmt.Println(countAndSay(Input))

	//Input = 4
	//Output = 1211
	Input = 4
	fmt.Println(countAndSay(Input))

	//Input = 5
	//Output = 111221
	Input = 5
	fmt.Println(countAndSay(Input))
}

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	var t strings.Builder
	ss, idx := "1", 0
	for i := 1; i < n; i++ {
		t.Reset()
		for j := 0; j < len(ss); j = idx {
			for idx = j; idx < len(ss) && ss[j] == ss[idx]; idx++ {
			}
			t.WriteString(strconv.Itoa(idx - j))
			t.WriteByte(ss[j])
		}
		ss = t.String()
	}

	return ss
}

/*
The count-and-say sequence is the sequence of integers with the first five terms as following:
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

    1. 1
    2. 11
    3. 21
    4. 1211
    5. 111221

  1 is read off as "one 1" or 11.
  1 被读作  "one 1"  ("一个一") , 即 11。

  11 is read off as "two 1s" or 21.
  11 被读作 "two 1s" ("两个一"）, 即 21。

  21 is read off as "one 2, then one 1" or 1211.
  21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

Given an integer n where 1 ≤ n ≤ 30, generate the n'th term of the count-and-say sequence.
给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

You can do so recursively, in other words from the previous member read off the digits, counting the number of digits in groups of the same digit.

Note: Each term of the sequence of integers will be represented as a string.
注意：整数序列中的每一项将表示为一个字符串。


Example 1:
示例 1:

  Input: 1
  输入: 1

  Output: "1"
  输出: "1"

  Explanation: This is the base case.
  解释：这是一个基本样例。


Example 2:
示例 2:

  Input: 4
  输入: 4

  Output: "1211"
  输出: "1211"

  Explanation: For n = 3 the term was "21" in which we have two groups "2" and "1", "2" can be read as "12" which means frequency = 1 and value = 2,
  解释：        当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2;
			   the same way "1" is read as "11", so the answer is the concatenation of "12" and "11" which is "1211".
			   类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
*/
