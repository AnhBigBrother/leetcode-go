package utils

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}
func (s *Stack[T]) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]
}
func (s *Stack[T]) Top() T {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}
	return s.items[len(s.items)-1]
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack[T]) Len() int {
	return len(s.items)
}
