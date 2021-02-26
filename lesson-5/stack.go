package main

type Stack struct {
	sli *List
}

func NewStack() *Stack {
	return &Stack{
		sli: &List{},
	}
}

func (s *Stack) Push(elem int) {
	node := &Node{
		Data: elem,
	}
	s.sli.Append(node)
}
func (s *Stack) Pop() int {
	if s.sli.Len() == 0 {
		return 0
	}
	elem := s.sli.Tail().Data
	s.sli.Delete(s.sli.Tail())
	return elem
}

//func main() {
//	stack := NewStack()
//	stack.Push(5)
//	stack.Push(6)
//	stack.Push(7)
//	fmt.Println(stack.Pop())
//	fmt.Println(stack.Pop())
//	fmt.Println(stack.Pop())
//}
