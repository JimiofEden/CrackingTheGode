package ch03

import "errors"

type QueueNode struct {
	data int
	next *QueueNode
}

func NewQueueNode(n int) *QueueNode {
	l := new(QueueNode)
	l.data = n
	return l
}

type Queue struct {
	first *QueueNode
	last *QueueNode
}

func (q *Queue) Add(n int) {
	newNode := NewQueueNode(n)
	newNode.next = q.last
	q.last = newNode
}
func (q *Queue) Remove() (int, error) {
	if q.first == nil {
		return 0, errors.New("Cannot remove from a nil queue")
	}
	v := q.first
	q.first = v.next
	return v.data, nil
}

func (q *Queue) Peek() (int, error) {
	if q.first == nil {
		return 0, errors.New("Cannot peek at a nil queue")
	}
	v := q.first
	return v.data, nil
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}