package collection

type List[T any] []T

func (l List[T]) ForEach(action func(T)) {
	for _, item := range l {
		action(item)
	}
}

func NewList[T any](elements ...T) List[T] {
	if len(elements) == 0 {
		return List[T](nil)
	}

	newList := make(List[T], 0, len(elements))
	for _, elem := range elements {
		newList = append(newList, elem)
	}

	return newList
}

func EmptyList[T any]() List[T] {
	return NewList[T]()
}
