package datastructures

/*
A basic stack, with "static" elements.
Elements can only be set at creation, afterwards, the stack is only used to keep track of the current item.
*/
type simpleStackConcrete[T comparable] struct {
	items        []T
	stackPointer int
}

type SimpleStack[T comparable] interface {
	Pop() (T, bool)
	Push()
	IsEmpty() bool
}

func NewSimpleStack[T comparable] (items Set[T]) SimpleStack[T] {
	return &simpleStackConcrete[T]{
		items: items.ToSlice(),
		stackPointer: items.Count() - 1,
	}
}

func (s *simpleStackConcrete[T]) Pop() (item T, exists bool) {
	// stack is already empty.
	if s.stackPointer == -1 {
		return item, false
	}

	s.stackPointer--
	return s.items[s.stackPointer + 1], true
}

func (s *simpleStackConcrete[T]) Push() {
	// stack is already at last value.
	if s.stackPointer == len(s.items) - 1 {
		return
	}
	s.stackPointer++
}

func (s *simpleStackConcrete[T]) IsEmpty() bool {
	return s.stackPointer == -1
}
