package main

import (
    "fmt"
)

func main() {
    Input := "III"
    //Output = 3
    fmt.Println(romanToInt(Input))

    Input = "IV"
    //Output = 4
    fmt.Println(romanToInt(Input))

    Input = "IX"
    //Output = 9
    fmt.Println(romanToInt(Input))

    Input = "LVIII"
    //Output = 58
    fmt.Println(romanToInt(Input))

    Input = "MCMXCIV"
    //Output = 1994
    fmt.Println(romanToInt(Input))
}

var romanToNumberMap = map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
    sum := 0
    pre := -1
    for i := len(s) - 1; i >= 0; i-- {
        current := romanToNumberMap[s[i]]
        if pre > current {
            sum -= current
            continue
        }
        sum += current
        pre = current
    }

    return sum
}

/*
func romanToInt1(s string) int {
    var res int

    for i := len(s) - 1; i >= 0; i-- {
        curRune := s[i]
        cur := romanToNumberMap[curRune]

        if i == 0 {
            res += cur
            break
        }

        if n := romanToNumberMap[s[i-1]]; n < cur {
            res += cur - n
            i--
        } else {
            res += cur
        }
    }

    return res
}
*/


/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

    Symbol       Value
    字符          数值
    I             1
    V             5
    X             10
    L             50
    C             100
    D             500
    M             1000

For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II.
The number twenty seven is written as XXVII, which is XX + V + II.
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the
number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number
nine, which is written as IX. There are six instances where subtraction is used:
通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到
的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

  I can be placed before V (5) and X (10) to make 4 and 9.
  I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。

  X can be placed before L (50) and C (100) to make 40 and 90.
  X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。

  C can be placed before D (500) and M (1000) to make 400 and 900. 
  C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

  Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
  给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。


Example 1:
示例 1:

  Input: "III"
  输入: "III"

  Output: 3
  输出: 3


Example 2:
示例 2:

  Input: "IV"
  输入: "IV"

  Output: 4
  输出: 4


Example 3:
示例 3:

  Input: "IX"
  输入: "IX"

  Output: 9
  输出: 9


Example 4:
示例 4:

  Input: "LVIII"
  输入: "LVIII"

  Output: 58
  输出: 58

  Explanation: L = 50, V= 5, III = 3.
  解释: L = 50, V= 5, III = 3.


Example 5:
示例 5:

  Input: "MCMXCIV"
  输入: "MCMXCIV"

  Output: 1994
  输出: 1994

  Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
  解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/

/* 解法二
func romanToInt(s string) int {
    var x int

    for i := len(s) - 1; i >= 0; i-- {
        switch s[i] {
        case 'I':
            x += 1
        case 'V':
            x += 5
            if i != 0 && s[i-1] == 'I' {
                x--
                i--
            }
        case 'X':
            x += 10
            if i != 0 && s[i-1] == 'I' {
                x--
                i--
            }
        case 'L':
            x += 50
            if i != 0 && s[i-1] == 'X' {
                x -= 10
                i--
            }
        case 'C':
            x += 100
            if i != 0 && s[i-1] == 'X' {
                x -= 10
                i--
            }
        case 'D':
            x += 500
            if i != 0 && s[i-1] == 'C' {
                x -= 100
                i--
            }
        case 'M':
            x += 1000
            if i != 0 && s[i-1] == 'C' {
                x -= 100
                i--
            }
        default:
            return 0
        }
    }

    return x
}
*/
