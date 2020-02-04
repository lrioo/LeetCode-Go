package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	//output := 2
	output := removeElement(nums, val)
	fmt.Println(output, nums[:output])

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	//output = 5
	output = removeElement(nums, val)
	fmt.Println(output, nums[:output])
}

func removeElement(nums []int, val int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}

/*
Given an array nums and a value val, remove all instances of that value in-place and return the new length.
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

The order of elements can be changed. It doesn't matter what you leave beyond the new length.
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。


Example 1:
示例 1:

  Given nums = [3,2,2,3], val = 3,
  给定 nums = [3,2,2,3], val = 3,

  Your function should return length = 2, with the first two elements of nums being 2.
  函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

  It doesn't matter what you leave beyond the returned length.
  你不需要考虑数组中超出新长度后面的元素。


Example 2:
示例 2:

  Given nums = [0,1,2,2,3,0,4,2], val = 2,
  给定 nums = [0,1,2,2,3,0,4,2], val = 2,

  Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.
  函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。

  Note that the order of those five elements can be arbitrary.
  注意这五个元素可为任意顺序。

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
    // nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
    int len = removeElement(nums, val);

    // any modification to nums in your function would be known by the caller.
    // 在函数里修改输入数组对于调用者是可见的。
    // using the length returned by your function, it prints the first len elements.
    // 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
    for (int i = 0; i < len; i++) {
        print(nums[i]);
    }
*/
