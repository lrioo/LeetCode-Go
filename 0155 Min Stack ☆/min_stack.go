package main

import "fmt"

func main() {
	actions := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}
	data := [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}
	output := test(actions, data)
	fmt.Println(output)
}

func test(actions []string, data [][]int) []interface{} {
	var ms MinStack
	output := make([]interface{}, len(actions))
	for i, action := range actions {
		d := data[i]

		switch action {
		case "MinStack":
			ms = Constructor()
		case "push":
			ms.Push(d[0])
		case "pop":
			ms.Pop()
		case "top":
			output[i] = ms.Top()
		case "getMin":
			output[i] = ms.GetMin()
		}
	}
	return output
}

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	if this.stack[len(this.stack)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
设计一个支持push，pop，top操作，并能在常数时间内检索到最小元素的栈。

push(x) -- Push element x onto stack.
        —— 将元素 x 推入栈中。
pop() -- Removes the element on top of the stack.
      —— 删除栈顶的元素。
top() -- Get the top element.
      —— 获取栈顶元素。
getMin() -- Retrieve the minimum element in the stack.
         —— 检索栈中的最小元素。


Example:
示例:

  Input:
  输入：
    ["MinStack","push","push","push","getMin","pop","top","getMin"]
    [[],[-2],[0],[-3],[],[],[],[]]

  Output:
  输出：
    [null,null,null,null,-3,null,0,-2]

  Explanation:
  解释：
    MinStack minStack = new MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    minStack.getMin(); // return -3  --> 返回 -3.
    minStack.pop();
    minStack.top();    // return 0   --> 返回 0.
    minStack.getMin(); // return -2  --> 返回 -2.


Constraints:
提示：

  Methods pop, top and getMin operations will always be called on non-empty stacks.
  pop、top和getMin操作总是在非空栈上调用。
*/
