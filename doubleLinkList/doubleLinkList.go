package doubleLinkList

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoubleLinkList struct {
	Length int
	Head   *Node
}

func NewDoubleLinkList() *DoubleLinkList {
	node := &Node{
		Prev: nil,
		Next: nil,
	}
	return &DoubleLinkList{
		Length: 0,
		Head:   node,
	}
}

func (l *DoubleLinkList) Append(e int) error {
	firstNode := l.Head.Next
	lastNode := l.Head.Prev

	// 先解决新节点的前驱和后继节点
	newNode := &Node{
		Data: e,
		Prev: lastNode,
		Next: firstNode,
	}

	if l.Length == 0 {
		l.Head.Prev = newNode
		l.Head.Next = newNode
	} else {
		lastNode.Next = newNode
		l.Head.Prev = newNode
	}

	l.Length++

	return nil
}

func (l *DoubleLinkList) Get(i int) (*Node, error) {
	if err := checkIndex(i); err != nil {
		return nil, err
	}

	// 最后一个节点
	if i == l.Length {
		return l.Head.Prev, nil
	}

	node := l.Head
	for j := 0; j < i; j++ {
		node = node.Next
	}

	return node, nil
}

func (l *DoubleLinkList) Insert(i int, e int) error {
	if err := checkIndex(i); err != nil {
		return err
	}

	newNode := &Node{
		Data: e,
	}

	if i == 1 {

	}
	node := l.Head
	for j := 0; j < i; j++ {
		node = node.Next
	}

	newNode.Prev = node.Prev
	newNode.Next = node
	node.Prev.Next = newNode
	node.Prev = newNode

	l.Length++

	return nil
}

func (l *DoubleLinkList) Delete(i int) error {
}

func (l *DoubleLinkList) checkIndex(i int) error {
	if i < 1 || i > l.Length {
		return fmt.Errorf("the index is invalid")
	}

	return nil
}
