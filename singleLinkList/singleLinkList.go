package singleLinkList

import "fmt"

type DataType int

type Node struct {
	Data DataType
	Next *Node
}

type SingleLinkList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func NewSingleLinkList() *SingleLinkList {
	node := &Node{
		Next: nil,
	}

	return &SingleLinkList{
		Length: 0,
		Head:   node,
		Tail:   node,
	}
}

func (l *SingleLinkList) Append(e DataType) (DataType, error) {
	node := &Node{
		Data: i,
	}

	l.Tail.Next = nil
	if l.Length == 0 {
	}

	l.Tail = node
}

func (l *SingleLinkList) Get(i int) (*Node, error) {
	if err := checkIndex(i); err != nil {
		return nil, err
	}

	node := l.Head
	for j := 1; j < i; j++ {
		node = node.Next
	}

	return node, nil

}

func (l *SingleLinkList) Insert(i int, e DataType) error {
	if err := checkIndex(i); err != nil {
		return err
	}

	newNode := &Node{
		Data: e,
	}

	node := l.Head
	for j := 1; j < i; j++ {
		node = node.Next
	}
	newNode.Next = node.Next
	node.Next = newNode

	l.Length += 1

	return nil
}

func (l *SingleLinkList) Delete(i int) (DataType, error) {

}

func (l *SingleLinkList) checkIndex(i int) error {
	if i < 1 || i > l.Length {
		return fmt.Errorf("the index is invalid")
	}
}
