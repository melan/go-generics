package collection

type Set[T comparable] map[T]struct{}

func (s Set[T]) ForEach(action func(T)) {
	for key := range s {
		action(key)
	}
}

func NewSet[T comparable](elements ...T) Set[T] {
	newSet := make(Set[T], len(elements))
	for _, elem := range elements {
		newSet[elem] = struct{}{}
	}

	return newSet
}

func EmptySet[T comparable]() Set[T] {
	return NewSet[T]()
}
