package ch02

type Node struct {
	value int
	next *Node
}

func NewNode(n int) *Node {
	l := new(Node)
	l.value = n
	return l
}

func (l *Node) AddTail(v int) {
	newItem := NewNode(v)
	for l.next != nil {
		l = l.next
	}
	l.next = newItem
}

func (l *Node) Len() int {
	i := 0
	for n := l; n != nil; n = n.next {
		i++
	}
	return i
}