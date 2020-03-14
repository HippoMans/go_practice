package Architecture

type Stack struct {
	link *LinkedList
}

func NewStack() *Stack {
	return &Stack{&LinkedList{}}
}

func (s *Stack) IsEmpty() bool {
	return s.link.IsEmpty()
}

func (s *Stack) Push(val int) {
	s.link.AddNode(val)
}

func (s *Stack) Pop() int {
	back_value := s.link.Back()
	s.link.PopBack()
	return back_value
}

func (s *Stack) PrintStack() {
	s.link.PrintNode()
}
