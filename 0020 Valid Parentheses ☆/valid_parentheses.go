package main

import (
    "fmt"
)

func main() {
    Input := "()"
    //Output = true
    fmt.Println(isValid(Input))

    Input = "()[]{}"
    //Output = true
    fmt.Println(isValid(Input))

    Input = "(]"
    //Output = false
    fmt.Println(isValid(Input))

    Input = "([)]"
    //Output = false
    fmt.Println(isValid(Input))

    Input = "{[]}"
    //Output = true
    fmt.Println(isValid(Input))
}

func isValid(str string) bool {
    if str == "" {
        return true
    }

    if len(str)%2 != 0 {
        return false
    }

    c := make([]int32, 0, len(str)/2)
    for i, s := range str {
        switch s {
        case '(', '[', '{':
            c = append(c, s)
        default:
            if i == 0 {
                return false
            }
            switch s {
            case ')':
                if c[len(c)-1] != '(' {
                    return false
                }
            case ']':
                if c[len(c)-1] != '[' {
                    return false
                }
            case '}':
                if c[len(c)-1] != '{' {
                    return false
                }
            }
            c = c[:len(c)-1]
        }
    }

    if len(c) > 0 {
        return false
    }

    return true
}

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。


An input string is valid if:
有效字符串需满足：

  Open brackets must be closed by the same type of brackets.
  左括号必须用相同类型的右括号闭合。

  Open brackets must be closed in the correct order.
  左括号必须以正确的顺序闭合。

  Note that an empty string is also considered valid.
  注意空字符串可被认为是有效字符串。


Example 1:
示例 1:

  Input: "()"
  输入: "()"

  Output: true
  输出: true


Example 2:
示例 2:

  Input: "()[]{}"
  输入: "()[]{}"

  Output: true
  输出: true


Example 3:
示例 3:

  Input: "(]"
  输入: "(]"

  Output: false
  输出: false


Example 4:
示例 4:

  Input: "([)]"
  输入: "([)]"

  Output: false
  输出: false


Example 5:
示例 5:

  Input: "{[]}"
  输入: "{[]}"

  Output: true
  输出: true
*/
