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

func (l *SingleLinkList) Get(i int) (DataType, error) {

}

func (l *SingleLinkList) Insert(i int, e DataType) error {
	node := &Node{
		Data: e,
	}

	if i < 1 || i > l.Length {
		return fmt.Errorf("the index is invalid")
	}

	for j := 1; j < i; j++ {

	}
}

func (l *SingleLinkList) Delete(i int) (DataType, error) {

}
