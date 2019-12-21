package main

import (
    "fmt"
)

func main() {
    haystack, needle := "hello", "ll"
    //Output := 2
    output := strStr(haystack, needle)
    fmt.Println(output)

    haystack, needle = "aaaaa", "bba"
    //Output = -1
    output = strStr(haystack, needle)
    fmt.Println(output)
}

func strStr(haystack string, needle string) int {
    if needle == "" || haystack == needle {
        return 0
    }

    if haystack == "" || len(haystack) <= len(needle) {
        return -1
    }

    for i := 0; i <= len(haystack)-len(needle); i++ {
        if haystack[i] == needle[0] {
            if haystack[i:i+len(needle)] == needle {
                return i
            }
        }
    }

    return -1
}

/*
Implement strStr().
实现 strStr() 函数。

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。


Example 1:
示例 1:

  Input: haystack = "hello", needle = "ll"
  输入: haystack = "hello", needle = "ll"

  Output: 2
  输出: 2


Example 2:
示例 2:

  Input: haystack = "aaaaa", needle = "bba"
  输入: haystack = "aaaaa", needle = "bba"

  Output: -1
  输出: -1


Clarification:
说明:

  What should we return when needle is an empty string? This is a great question to ask during an interview.
  当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

  For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
  对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/
