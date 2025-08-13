package datastructures

import (
	"iter"
	"slices"
)

type setConcrete[T comparable] map[T]struct{}

type Set[T comparable] interface {
	Contains(key T) bool
	Add(key T)
	Remove(key T)
	Count() int
	All() iter.Seq[T]
	ToSlice() []T
}

func NewSet[T comparable]() Set[T] {
	return setConcrete[T]{}
}

func (s setConcrete[T]) Contains(key T) bool {
	_, exists := s[key]
	return exists
}

func (s setConcrete[T]) Add(key T) {
	if s.Contains(key) {
		return
	}
	s[key] = struct{}{}
}

func (s setConcrete[T]) Remove(key T) {
	delete(s, key)
}

func (s setConcrete[T]) Count() int {
	return len(s)
}

func (s setConcrete[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}

func (s setConcrete[T]) ToSlice() []T {
	return slices.Collect(s.All())
}
