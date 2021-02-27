package main

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	Data int
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

// 2 elements
//		head - next->
//->	prev - tail

// 3 elements and more
// 	  	head - next->
// -> 	prev - node - next->
// ->		   prev - tail
func (l *List) Add(position, node *Node) {
	l.len++
	// add first element
	//		   		head(nodeN) - next->
	//	(add)node - next
	if position == nil {
		node.next = l.head
		l.head = node
		return
	}
	// 1 element
	//		head ->
	//->	tail
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	// 	add node prev position
	//	prev - 	position - 				next->
	//				prev - (add)node - 	next
	node.prev = position
	node.next = position.next
	position.next = node
	if position == l.tail {
		l.tail = node
	}
}

func (l *List) Append(node *Node) {
	l.Add(l.tail, node)
	l.tail = node
}

func (l *List) Find(num int) *Node {
	for tmp := l.head; tmp != nil; tmp = tmp.next {
		if tmp.Data == num {
			return tmp
		}
	}
	return nil
}

func (l *List) Delete(node *Node) {
	if l.len == 1 {
		l.head = nil
		l.tail = nil
	}
	if l.head == node {
		l.head = l.head.next
	}
	if l.head != nil {
		for tmp := l.head; tmp != nil; tmp = tmp.next {
			if tmp.next == node && tmp.next != l.tail {
				tmp.next = node.next
			}
			if tmp.next == node && tmp.next == l.tail {
				l.tail = tmp
				tmp.next = nil
			}
		}
	}
	l.len--
}

func (l *List) Print() {
	if l.head != nil {
		i := 0
		for tmp := l.head; tmp != nil; tmp = tmp.next {
			fmt.Printf("[%d] - %d\n", i, tmp.Data)
			i++
		}
	}
}

func main() {
	list := &List{}
	emp := list.Find(3)
	fmt.Println(emp)

	for i := 0; i < 5; i++ {
		node := &Node{
			Data: i,
		}
		list.Append(node)
	}

	node1 := &Node{
		Data: 6,
	}
	list.Add(nil, node1)

	fmt.Println(list.Len())
	list.Print()

	node := list.Find(3)

	list.Delete(node)

	fmt.Println(list.Len())
	list.Print()
}
