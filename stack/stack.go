package stack

import "fmt"

type stack struct {
	elements []int
}

func (s *stack) Push(element int) {
	s.elements = append(s.elements, element)
	fmt.Println("Stack looks like:", s.elements)
	return
}
func (s *stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	poppedElement := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	fmt.Println("Stack looks like:", s.elements)
	return poppedElement, true
}
func BuildStack() stack {
	return stack{}
}
