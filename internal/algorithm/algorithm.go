package algorithm

type Set[T comparable] map[T]struct{}

func (s Set[T]) Contains(key T) bool {
	_, exists := s[key]
	return exists
}

func (s *Set[T]) Add(key T) {
	if s.Contains(key) {
		return
	}
	(*s)[key] = struct{}{}
}

func (s *Set[T]) Remove(key T) {
	if !s.Contains(key) {
		return
	}

	delete(*s, key)
}
