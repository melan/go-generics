package collection

type List[T any] struct {
	traversable[T]
	list []T
}

func (l *List[T]) ForEach(action func(T)) {
	for _, item := range l.list {
		action(item)
	}
}

func NewList[T any](elements ...T) *List[T] {
	list := &List[T]{}
	list.traversable.forEach = list

	if len(elements) == 0 {
		return list
	}

	list.list = make([]T, 0, len(elements))
	list.list = append(list.list, elements...)

	return list
}

func EmptyList[T any]() *List[T] {
	return NewList[T]()
}

func (l *List[T]) Add(elements ...T) *List[T] {
	l.list = append(l.list, elements...)

	return l
}
