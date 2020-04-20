package util

// Queue implemnents the queue
type Queue []Node

// IsEmpty check if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push on to stack
func (q *Queue) Push(node Node) {
	*q = append(*q, node)
}

// Pop on to stack
func (q *Queue) Pop() (Node, bool) {
	if q.IsEmpty() {
		return Node{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true

}
