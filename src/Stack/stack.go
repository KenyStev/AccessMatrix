package Stack

import (
	"fmt"
  "../Domain"
)

type Node struct {
	Value *Domain.Domain
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value.Name)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

func (s *Stack) GetCount() int{
  return s.count;
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Top() *Domain.Domain{
  if (s.count>=0){
    return s.nodes[s.count-1].Value
  }
  return nil
}
