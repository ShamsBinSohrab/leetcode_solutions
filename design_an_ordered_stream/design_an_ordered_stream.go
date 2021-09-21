package main

//Problem: https://leetcode.com/problems/design-an-ordered-stream/

func main() {
	str := Constructor(5)
	str.Insert(3, "c")
	str.Insert(1, "a")
	str.Insert(2, "b")
	str.Insert(5, "e")
	str.Insert(4, "d")

}

type OrderedStream struct {
	array []string
	ptr   int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{array: make([]string, n+1), ptr: 1}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	os.array[idKey] = value
	oldPtr := os.ptr
	if os.ptr == idKey {
		for os.ptr < len(os.array) && os.array[os.ptr] != "" {
			os.ptr++
		}
	}
	return os.array[oldPtr:os.ptr]
}
