package main

import "fmt"

func main() {
	fmt.Println("----------stack-----------")
	stack := &Stack{
		top: -1,
	}
	stack.Push(12)
	stack.Push(13)
	stack.Pop()
	stack.Push(14)
	fmt.Println(stack)

	fmt.Println("----------queue array-----------")
	queue := &QueueArray{}
	queue.EnQueue(33)
	queue.EnQueue(34)
	queue.DeQueue()
	queue.EnQueue(35)
	fmt.Println(queue)

	fmt.Println("----------queue slice-----------")
	queue2 := &QueueSlice{}
	queue2.EnQueue(33)
	queue2.EnQueue(34)
	queue2.DeQueue()
	queue2.EnQueue(35)
	fmt.Println(queue2)

	fmt.Println("----------queue stack-----------")
	qs := &QueueStack{}
	qs.Push(1)
	qs.Push(2)
	qs.Pop()
	qs.Push(3)
	qs.Push(4)
	qs.Pop()
	qs.Push(5)
	fmt.Println(qs)

	fmt.Println("----------slice stack-----------")
	ss := &SliceStack{}
	ss.Push(1)
	ss.Push(2)
	ss.Pop()
	ss.Push(3)
	ss.Push(4)
	ss.Pop()
	ss.Push(5)
	fmt.Println(ss)
}
