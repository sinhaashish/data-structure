package util

// Stack is the stack implementation
type Stack []Node

// IsEmpty  check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push on to stack
func (s *Stack) Push(node Node) {
	*s = append(*s, node)
}

// Pop on to stack
func (s *Stack) Pop() (Node, bool) {
	if s.IsEmpty() {
		return Node{}, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

// Peek on to stack
func (s *Stack) Peek() (Node, bool) {
	if s.IsEmpty() {
		return Node{}, false
	}
	element := (*s)[0]
	return element, true
}
