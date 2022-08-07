package collection

import (
	"golang.org/x/exp/constraints"
)

type Traversable[T any] interface {
	ForEach(action func(T))
}

func Map[T, V any](t Traversable[T], transform func(T) V) Traversable[V] {
	newList := EmptyList[V]()
	t.ForEach(func(elem T) {
		newList = append(newList, transform(elem))
	})

	return newList
}

func Head[T any](t Traversable[T]) (T, bool) {
	var found bool
	var first T

	t.ForEach(func(elem T) {
		if found {
			return
		}

		first = elem
		found = true
	})

	return first, found
}

func HeadOption[T any](t Traversable[T]) Option[T] {
	first, found := Head(t)

	if found {
		return Some(first)
	}

	return None[T]()
}

func Last[T any](t Traversable[T]) (T, bool) {
	var last T
	var found bool

	t.ForEach(func(elem T) {
		last = elem
		found = true
	})

	return last, found
}

func LastOption[T any](t Traversable[T]) Option[T] {
	last, found := Last(t)

	if found {
		return Some(last)
	}

	return None[T]()
}

func Find[T any](t Traversable[T], filter func(elem T) bool) Option[T] {
	var found bool
	var result Option[T] = None[T]()

	t.ForEach(func(elem T) {
		if found {
			return
		}

		if filter(elem) {
			result = Some(elem)
			found = true
		}
	})

	return result
}

func Exists[T any](t Traversable[T], filter func(elem T) bool) bool {
	return IsDefined[T](Find(t, filter))
}

func Count[T any](t Traversable[T], filter func(elem T) bool) int {
	var result int
	t.ForEach(func(elem T) {
		if filter == nil || filter(elem) {
			result++
		}
	})

	return result
}

func ForAll[T any](t Traversable[T], filter func(elem T) bool, action func(elem T)) {
	t.ForEach(func(elem T) {
		if filter == nil || filter(elem) {
			action(elem)
		}
	})
}

func IsEmpty[T any](t Traversable[T]) bool {
	return Count[T](HeadOption(t), nil) == 0
}

func NonEmpty[T any](t Traversable[T]) bool {
	return !IsEmpty(t)
}

func IsDefined[T any](t Traversable[T]) bool {
	return NonEmpty(t)
}

func ToSet[T comparable](t Traversable[T]) Set[T] {
	switch t1 := t.(type) {
	case Set[T]:
		return t1

	default:
		newSet := EmptySet[T]()
		t.ForEach(func(elem T) {
			newSet[elem] = struct{}{}
		})

		return newSet
	}
}

func ToList[T any](t Traversable[T]) List[T] {
	switch t1 := t.(type) {
	case List[T]:
		return t1

	default:
		newList := EmptyList[T]()

		t.ForEach(func(elem T) {
			newList = append(newList, elem)
		})

		return newList
	}
}

type number interface {
	constraints.Integer | constraints.Float
}

func Sum[T number](t Traversable[T]) T {
	var result T
	t.ForEach(func(elem T) {
		result += elem
	})

	return result
}

func Product[T number](t Traversable[T]) T {
	var result T
	t.ForEach(func(elem T) {
		result *= elem
	})

	return result
}

func Min[T constraints.Ordered](t Traversable[T]) T {
	var initialized bool
	var result T
	t.ForEach(func(elem T) {
		if !initialized {
			initialized = true
			result = elem
			return
		}

		if result > elem {
			result = elem
		}
	})

	return result
}

func Max[T constraints.Ordered](t Traversable[T]) T {
	var initialized bool
	var result T
	t.ForEach(func(elem T) {
		if !initialized {
			initialized = true
			result = elem
			return
		}

		if result < elem {
			result = elem
		}
	})

	return result
}

func Fold[T any](t Traversable[T], initialValue T, op func(a, b T) T) T {
	accumulator := initialValue

	t.ForEach(func(elem T) {
		accumulator = op(accumulator, elem)
	})

	return accumulator
}

func Reduce[T any](t Traversable[T], op func(a, b T) T) T {
	var initialized bool
	var accumulator T

	t.ForEach(func(elem T) {
		if !initialized {
			accumulator = elem
			initialized = true
			return
		}

		accumulator = op(accumulator, elem)
	})

	return accumulator
}
