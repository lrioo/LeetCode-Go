package main

import "fmt"

func main() {
	nums1, m := []int{1, 2, 3, 0, 0, 0}, 3
	nums2, n := []int{2, 5, 6}, 3
	//Output: [1,2,2,3,5,6]
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		}

		if nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
给你两个有序整数数组nums1和nums2，请你将nums2合并到nums1中，使nums1成为一个有序数组。


Note:
说明:

  The number of elements initialized in nums1 and nums2 are m and n respectively.
  初始化nums1和nums2的元素数量分别为m和n。

  You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
  你可以假设nums1有足够的空间（空间大小大于或等于m+n）来保存nums2中的元素。


Example:
示例:

  Input:
  输入:
    nums1 = [1,2,3,0,0,0], m = 3
    nums2 = [2,5,6], n = 3

  Output: [1,2,2,3,5,6]
  输出: [1,2,2,3,5,6]
*/
