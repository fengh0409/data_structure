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
}

func NewSingleLinkList() *SingleLinkList {
	node := &Node{
		Next: nil,
	}

	return &SingleLinkList{
		Length: 0,
		Head:   node,
	}
}

func (l *SingleLinkList) Append(e DataType) error {
	newNode := &Node{
		Data: e,
	}

	node := l.Head
	for j := 0; j < l.Length; j++ {
		node = node.Next
	}
	if node.Next == nil {
		node.Next = newNode
	}

	l.Length++

	return nil
}

func (l *SingleLinkList) Get(i int) (*Node, error) {
	if err := checkIndex(i); err != nil {
		return nil, err
	}

	node := l.Head
	for j := 0; j < i; j++ {
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
	for j := 0; j < i-1; j++ {
		node = node.Next
	}
	newNode.Next = node.Next
	node.Next = newNode

	l.Length++

	return nil
}

func (l *SingleLinkList) Delete(i int) (DataType, error) {

}

func (l *SingleLinkList) checkIndex(i int) error {
	if i < 1 || i > l.Length {
		return fmt.Errorf("the index is invalid")
	}
}
