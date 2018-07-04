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

func (l *DoubleLinkList) Get(i int) {
}

func (l *DoubleLinkList) GetLast(i int) {
}

func (l *DoubleLinkList) Insert(i int, e int) {
}

func (l *DoubleLinkList) Delete(i int) {
}
