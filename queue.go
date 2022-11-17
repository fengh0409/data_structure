package main

import "fmt"

const maxQueueSize = 10

// 线性队列，数组实现
type QueueArray struct {
	data   [maxQueueSize]int
	length int
}

func (q *QueueArray) EnQueue(e int) error {
	newLength := q.length + 1
	if newLength == maxQueueSize {
		return fmt.Errorf("stack overflow")
	}
	q.data[q.length] = e
	q.length = newLength
	return nil
}

func (q *QueueArray) DeQueue() (int, error) {
	if q.length == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	data := q.data[0]
	for i := 0; i < q.length-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.data[q.length] = 0
	q.length--
	if q.length == 0 {
		q.data[0] = 0
	}

	return data, nil
}

func (q *QueueArray) Length() int {
	return q.length
}

// 两个队列实现一个栈，QueueArray实现
type QueueStack struct {
	queue1 QueueArray
	queue2 QueueArray
}

func (qs *QueueStack) Pop() (int, error) {
	length1 := qs.queue1.Length()
	if length1 > 0 {
		for i := 0; i < length1-1; i++ {
			data, _ := qs.queue1.DeQueue()
			qs.queue2.EnQueue(data)
		}
		return qs.queue1.DeQueue()
	}

	length2 := qs.queue2.Length()
	if length2 > 0 {
		for i := 0; i < length2-1; i++ {
			data, _ := qs.queue2.DeQueue()
			qs.queue1.EnQueue(data)
		}
		return qs.queue2.DeQueue()
	}
	return 0, fmt.Errorf("queue1 and queue2 are both empty")
}

func (qs *QueueStack) Push(e int) error {
	var err error
	if qs.queue1.Length() > 0 {
		err = qs.queue1.EnQueue(e)
	} else {
		err = qs.queue2.EnQueue(e)
	}
	return err
}

// 线性队列，slice实现
type QueueSlice struct {
	data   []int
	length int
}

func (q *QueueSlice) EnQueue(e int) error {
	q.data = append(q.data, e)
	q.length++
	return nil
}

func (q *QueueSlice) DeQueue() (int, error) {
	if q.length == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	data := q.data[0]
	q.data = q.data[1:]
	q.length--
	return data, nil
}

// 两个队列实现一个栈，Slice实现
type SliceStack struct {
	queue1 []int
	queue2 []int
}

func (ss *SliceStack) Pop() (int, error) {
	if len(ss.queue1) > 0 {
		for i := 0; i < len(ss.queue1)-1; i++ {
			ss.queue2 = append(ss.queue2, ss.queue1[i])
		}
		data := ss.queue1[len(ss.queue1)-1]
		ss.queue1 = make([]int, 0)
		return data, nil
	}
	if len(ss.queue2) > 0 {
		for i := 0; i < len(ss.queue2)-1; i++ {
			ss.queue1 = append(ss.queue1, ss.queue2[i])
		}
		data := ss.queue2[len(ss.queue2)-1]
		ss.queue2 = make([]int, 0)
		return data, nil
	}
	return 0, fmt.Errorf("slice stack is empty")
}

func (ss *SliceStack) Push(e int) error {
	if len(ss.queue1) > 0 {
		ss.queue1 = append(ss.queue1, e)
	} else {
		ss.queue2 = append(ss.queue2, e)
	}
	return nil
}

// 队列，链表实现
type QueueNode struct {
	data int
	next *QueueNode
}

type QueueList struct {
	rear *QueueNode
}

func (q *QueueList) EnQueue(e int) error {
	node := &QueueNode{
		data: e,
	}
	if q.rear != nil {
		node.next = q.rear
	}
	q.rear = node
	return nil
}

func (q *QueueList) DeQueue() (int, error) {
	if q.rear == nil {
		return 0, fmt.Errorf("queue is empty")
	}

	return 0, nil
}

// 两个队列实现一个栈，链表实现
type QueueListStack struct {
	queue1 *QueueList
	queue2 *QueueList
}

func (qls *QueueListStack) Pop() {

}

func (qls *QueueListStack) Push() {

}
