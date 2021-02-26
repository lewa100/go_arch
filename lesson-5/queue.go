package main

import "fmt"

type Queue struct {
	sli *List
}

func NewQueue() *Queue {
	return &Queue{
		sli: &List{},
	}
}

func (s *Queue) Push(elem int) {
	node := &Node{
		Data: elem,
	}
	s.sli.Append(node)
}
func (s *Queue) Pop() int {
	if s.sli.Len() == 0 {
		return 0
	}
	elem := s.sli.Head().Data
	s.sli.Delete(s.sli.Head())
	return elem
}

func main() {
	queue := NewQueue()
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
