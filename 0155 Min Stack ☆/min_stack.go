package main

func main() {

}

type MinStack struct {
}

/** initialize your data structure here. */
func Constructor() MinStack {

}

func (this *MinStack) Push(x int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {

}

func (this *MinStack) GetMin() int {

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
