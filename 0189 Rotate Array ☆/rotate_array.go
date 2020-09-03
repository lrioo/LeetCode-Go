package main

func main() {

}

func rotate(nums []int, k int) {

}

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.
给定一个数组，将数组中的元素向右移动k个位置，其中k是非负数。


Example 1:
示例 1：

  Input: nums = [1,2,3,4,5,6,7], k = 3
  输入：[1,2,3,4,5,6,7]，k = 3

  Output: [5,6,7,1,2,3,4]
  输出：[5,6,7,1,2,3,4]

  Explanation:
  解释：

    rotate 1 steps to the right: [7,1,2,3,4,5,6]
    向右旋转1步：[7,1,2,3,4,5,6]

    rotate 2 steps to the right: [6,7,1,2,3,4,5]
    向右旋转2步：[6,7,1,2,3,4,5]

    rotate 3 steps to the right: [5,6,7,1,2,3,4]
    向右旋转3步：[5,6,7,1,2,3,4]


Example 2:
示例2：

  Input: nums = [-1,-100,3,99], k = 2
  输入：[-1,-100,3,99]，k = 2

  Output: [3,99,-1,-100]
  输出：[3,99,-1,-100]

  Explanation:
  解释：

    rotate 1 steps to the right: [99,-1,-100,3]
    向右旋转1步：[99,-1,-100,3]

    rotate 2 steps to the right: [3,99,-1,-100]
    向右旋转2步：[3,99,-1,-100]


Constraints:
说明：

  1. 1 <= nums.length <= 2 * 10^4
  1. nums的长度为 1 ~ 2*10^4。

  2. It's guaranteed that nums[i] fits in a 32 bit-signed integer.
  2. 保证nums的任意元素为32位整型。

  3. k >= 0
  3. k>=0


Follow up:
进阶：

  Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
  尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。

  Could you do it in-place with O(1) extra space?
  要求使用空间复杂度为O(1)的原地算法。
*/
