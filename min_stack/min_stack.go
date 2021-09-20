package main

func main() {
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	minstack.GetMin()
	minstack.Pop()
	minstack.Top()
	minstack.GetMin()

}

type MinStack struct {
	list []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.list = append(ms.list, val)
	if len(ms.min) == 0 || ms.GetMin() >= val {
		ms.min = append(ms.min, val)
	}
}

func (ms *MinStack) Pop() {
	listTop := len(ms.list) - 1
	minTop := len(ms.min) - 1
	if ms.list[listTop] == ms.min[minTop] {
		ms.min = ms.min[:minTop]
	}
	ms.list = ms.list[:listTop]
}

func (ms *MinStack) Top() int {
	return ms.list[len(ms.list)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min[len(ms.min)-1]
}
