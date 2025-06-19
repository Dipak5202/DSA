package stack

type Stack struct {
	Top      int
	Elements [5]int
}

func CreateStack() *Stack {
	return &Stack{Top: -1}
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) IsFull() bool {
	return s.Top == len(s.Elements)-1
}

func (s *Stack) Push(value int) bool {
	if s.IsFull() {
		return false
	}

	s.Top++
	s.Elements[s.Top] = value
	return true
}

func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}

	println("Popped value:", s.Elements[s.Top])
	s.Top--
	return true
}

func (s *Stack) Peek() bool {
	if s.IsEmpty() {
		return false
	}
	println("Peeked value:", s.Elements[s.Top])
	return true
}
