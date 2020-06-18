package main

import "fmt"

func main() {
	Input := 0
	//Output := [1,3,3,1]
	fmt.Println(getRow(Input))
}

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}

/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
给定一个非负索引k，其中k≤33，返回杨辉三角的第k行。

Note that the row index starts from 0.
注意，行的索引从0开始。

In Pascal's triangle, each number is the sum of the two numbers directly above it.
在杨辉三角中，每个数是它左上方和右上方的数的和。


Example:
示例:

  Input: 3
  输入: 3

  Output: [1,3,3,1]
  输出: [1,3,3,1]


Follow up:
进阶：

  Could you optimize your algorithm to use only O(k) extra space?
  你可以优化你的算法到O(k)空间复杂度吗？
*/
