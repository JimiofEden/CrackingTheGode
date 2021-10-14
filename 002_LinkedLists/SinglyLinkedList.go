package ch02

type SinglyLinkedList struct {
	value int
	next *SinglyLinkedList
}

func NewSinglyLinkedList(n int) *SinglyLinkedList {
	l := new(SinglyLinkedList)
	l.value = n
	return l
}

func (l *SinglyLinkedList) AddTail(v int) {
	newItem := NewSinglyLinkedList(v)
	for l.next != nil {
		l = l.next
	}
	l.next = newItem
}

func (l *SinglyLinkedList) Len() int {
	i := 0
	for n := l; n != nil; n = n.next {
		i++
	}
	return i
}