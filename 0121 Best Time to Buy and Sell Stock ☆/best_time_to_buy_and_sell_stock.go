package main

import (
	"fmt"
)

func main() {
	Input := []int{7, 1, 5, 3, 6, 4}
	//Output := 5
	fmt.Println(maxProfit(Input))

	Input = []int{7, 6, 4, 3, 1}
	//Output := 0
	fmt.Println(maxProfit(Input))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	pro, cost := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if cost > prices[i] {
			cost = prices[i]
		}

		if diff := prices[i] - cost; pro < diff {
			pro = diff
		}
	}
	return pro
}

/*
Say you have an array for which the ith element is the price of a given stock on day i.
给定一个数组，它的第i个元素是一支给定股票第i天的价格。

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),
design an algorithm to find the maximum profit.
如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

Note that you cannot sell a stock before you buy one.
注意：你不能在买入股票前卖出股票。


Example 1:
示例 1:

  Input: [7,1,5,3,6,4]
  输入: [7,1,5,3,6,4]

  Output: 5
  输出: 5

  Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
               Not 7-1 = 6, as selling price needs to be larger than buying price.
  解释: 在第2天（股票价格=1）的时候买入，在第5天（股票价格=6）的时候卖出，最大利润=6-1=5。
       注意利润不能是7-1=6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。


Example 2:
示例 2:

  Input: [7,6,4,3,1]
  输入: [7,6,4,3,1]

  Output: 0
  输出: 0

  Explanation: In this case, no transaction is done, i.e. max profit = 0.
  解释: 在这种情况下, 没有交易完成, 所以最大利润为0。
*/
