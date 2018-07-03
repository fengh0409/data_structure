package doubleLinkList

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
	return &DoubleLinkList{
		Length: 0,
		Head:   nil,
	}
}

func (l *DoubleLinkList) Append(e int) error {
	var prev, next *Node
	if l.Length == 0 {
		prev = nil
		next = nil
	} else {
		prev = l.Head.Prev
		next = l.Head.Next
	}

	node := &Node{
		Data: e,
		Prev: prev,
		Next: next,
	}
	l.Head.Prev = node
	l.Length++

	return nil
}

func (l *DoubleLinkList) Get(i int) {
}

func (l *DoubleLinkList) GetLast(i int) {
}

func (l *DoubleLinkList) Insert(i int, e int) {
}

func (l *DoubleLinkList) Delete(i int) {
}
