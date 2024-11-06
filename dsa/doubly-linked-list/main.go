package main

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func newLinkedList() *linkedList {
	return new(linkedList)
}

func (l *linkedList) prepend(n *node) *linkedList {
	if l.length == 0 {
		return l.append(n)
	}
	n.next, l.head = l.head, n
	l.length++
	return l
}
func (l *linkedList) insertAt() {}
func (l *linkedList) append(n *node) *linkedList {
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		l.tail.next, l.tail = n, n
	}
	l.length++
	return l
}

func (l *linkedList) remove(idx int) (*int, error) {
	if idx >= l.length || idx < 0 {
		return nil, errors.New("out of bounds")
	}
	if idx == 1 {
		temp := l.head.next
		l.head, l.tail = nil, nil
		return &temp.value, nil
	}
	return recurseDelete(l.head, l.head.next, idx-1), nil
}

func (l *linkedList) get(idx int) (*node, error) {
	if idx >= l.length || idx < 0 {
		return nil, errors.New("out of bounds")
	}
	if idx == l.length-1 {
		return l.tail, nil
	}
	return recurse(l.head, idx), nil
}

func recurse(n *node, idx int) *node {
	if idx == 0 {
		return n
	}
	idx--
	return recurse(n.next, idx)
}

func recurseDelete(prev, present *node, idx int) *int {
	if idx == 0 {
		prev.next = present.next

		return &present.value
	}
	idx--
	return recurseDelete(present, present.next, idx)
}

func main() {
	fmt.Println(newLinkedList())
}
