package stack

/*
Design a stack class that supports the push, pop, top, and getMin operations.

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
Each function should run in O(1) time.

Example 1:

Input: ["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]

Output: [null,null,null,null,0,null,2,1]

Explanation:
MinStack minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(0);
minStack.getMin(); // return 0
minStack.pop();
minStack.top();    // return 2
minStack.getMin(); // return 1
Constraints:

-2^31 <= val <= 2^31 - 1.
pop, top and getMin will always be called on non-empty stacks.

*/

type MinStack struct {
	stack     []int
	monoStack []int // min mono stack
}

func Constructor() MinStack {
	return MinStack{
		stack:     []int{},
		monoStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	if len(s.monoStack) == 0 || s.monoStack[len(s.monoStack)-1] >= val { // we don't care about larger values since they are guaranteed to be popped before elements already in mono stack
		s.monoStack = append(s.monoStack, val)
	}
	s.stack = append(s.stack, val)
}

func (s *MinStack) Pop() {
	popped := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if popped == s.monoStack[len(s.monoStack)-1] {
		s.monoStack = s.monoStack[:len(s.monoStack)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.monoStack[len(s.monoStack)-1]
}
