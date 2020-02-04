package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	//output := 2
	output := removeDuplicates(nums)
	fmt.Println(output, nums[:output])

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//output = 5
	output = removeDuplicates(nums)
	fmt.Println(output, nums[:output])
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[idx] != nums[i] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}

/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。


Example 1:
示例 1:

  Given nums = [1,1,2],
  给定数组 nums = [1,1,2],

  Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
  函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

  It doesn't matter what you leave beyond the returned length.
  你不需要考虑数组中超出新长度后面的元素。


Example 2:
示例 2:

  Given nums = [0,0,1,1,1,2,2,3,3,4],
  给定 nums = [0,0,1,1,1,2,2,3,3,4],

  Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
  函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

  It doesn't matter what values are set beyond the returned length.
  你不需要考虑数组中超出新长度后面的元素。


Clarification:
说明:

  Confused why the returned value is an integer but your answer is an array?
  为什么返回数值是整数，但输出的答案是数组呢?

  Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.
  请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

  Internally you can think of this:
  你可以想象内部操作如下:
    // nums is passed in by reference. (i.e., without making a copy)
    // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
    int len = removeDuplicates(nums);

    // any modification to nums in your function would be known by the caller.
    // 在函数里修改输入数组对于调用者是可见的。
    // using the length returned by your function, it prints the first len elements.
    // 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
    for (int i = 0; i < len; i++) {
        print(nums[i]);
    }
*/
