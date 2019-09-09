package main

import (
    "fmt"
)

func main() {
    Input := []string{"flower", "flow", "flowll"}
    //Output = "fl"
    fmt.Println(longestCommonPrefix(Input))

    Input = []string{"dog", "racecar", "car"}
    //Output = ""
    fmt.Println(longestCommonPrefix(Input))

    Input = []string{"aaa", "aa", "a"}
    //Output = ""
    fmt.Println(longestCommonPrefix(Input))
}

func longestCommonPrefix(strs []string) string {
    var prefix string

    for i := 0; i < len(strs); i++ {
        if i == 0 {
            prefix = strs[0]
            continue
        }

        var j int
        for j = 0; j < len(prefix) && j < len(strs[i]); j++ {
            if prefix[j] != strs[i][j] {
                if j == 0 {
                    return ""
                }
                break
            }
        }
        prefix = prefix[:j]
    }

    return prefix
}

/*
Write a function to find the longest common prefix string amongst an array of strings.
编写一个函数来查找字符串数组中的最长公共前缀。

If there is no common prefix, return an empty string "".
如果不存在公共前缀，返回空字符串 ""。


Example 1:
示例 1:

  Input: ["flower","flow","flight"]
  输入: ["flower","flow","flight"]

  Output: "fl"
  输出: "fl"


Example 2:
示例 2:

  Input: ["dog","racecar","car"]
  输入: ["dog","racecar","car"]

  Output: ""
  输出: ""
  Explanation: There is no common prefix among the input strings.
  解释: 输入不存在公共前缀。


Note:
说明:

  All given inputs are in lowercase letters a-z.
  所有输入只包含小写字母 a-z 。
*/
