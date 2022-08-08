package collection

type Set[T comparable] struct {
	traversable[T]
	set map[T]struct{}
}

func (s Set[T]) ForEach(action func(T)) {
	for key := range s.set {
		action(key)
	}
}

func NewSet[T comparable](elements ...T) *Set[T] {
	newSet := &Set[T]{
		set: make(map[T]struct{}, len(elements)),
	}
	newSet.traversable.forEach = newSet
	for _, elem := range elements {
		newSet.set[elem] = struct{}{}
	}

	return newSet
}

func EmptySet[T comparable]() *Set[T] {
	return NewSet[T]()
}

func (s *Set[T]) Add(elements ...T) *Set[T] {
	for _, elem := range elements {
		s.set[elem] = struct{}{}
	}

	return s
}
