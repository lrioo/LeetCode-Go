package main

import (
	. "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	Input := 5
	fmt.Println(TwoDimensionalArray(Slice(generate(Input))))
}

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	rows := make([][]int, numRows)
	rows[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i] = row
	}
	return rows
}

func generateR(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	return generateRecursive(numRows)
}

func generateRecursive(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	rows := generateRecursive(numRows - 1)
	row := make([]int, numRows)
	row[0], row[numRows-1] = 1, 1
	for j := 1; j < numRows-1; j++ {
		row[j] = rows[len(rows)-1][j-1] + rows[len(rows)-1][j]
	}

	return append(rows, row)
}

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

In Pascal's triangle, each number is the sum of the two numbers directly above it.
在杨辉三角中，每个数是它左上方和右上方的数的和。

Example:
示例:

  Input: 5
  输入: 5

  Output:
  输出:

      [
           [1],
          [1,1],
         [1,2,1],
        [1,3,3,1],
       [1,4,6,4,1]
     ]
*/
