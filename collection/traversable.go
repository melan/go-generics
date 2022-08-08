package collection

import (
	"golang.org/x/exp/constraints"
)

type forEach[T any] interface {
	ForEach(action func(T))
}

type Traversable[T any] interface {
	forEach[T]
	Head() (T, bool)
	HeadOption() *Option[T]
	GetOrElse(altValue T) T
	Last() (T, bool)
	LastOption() *Option[T]
	Find(filter func(elem T) bool) *Option[T]
	Exists(filter func(elem T) bool) bool
	Count(filter func(elem T) bool) int
	Select(filter func(elem T) bool) *List[T]
	ForAll(filter func(elem T) bool, transform func(elem T) T) *List[T]
	IsEmpty() bool
	NonEmpty() bool
	IsDefined() bool
	Reduce(op func(a, b T) T) T
	ToList() *List[T]
}

type traversable[T any] struct {
	forEach[T]
}

var _ Traversable[string] = &traversable[string]{}

func (t *traversable[T]) Head() (T, bool) {
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

func (t *traversable[T]) HeadOption() *Option[T] {
	first, found := t.Head()

	if found {
		return Some(first)
	}

	return None[T]()
}

func (t *traversable[T]) GetOrElse(altValue T) T {
	head, found := t.Head()
	if found {
		return head
	}

	return altValue
}

func (t *traversable[T]) Last() (T, bool) {
	var last T
	var found bool

	t.ForEach(func(elem T) {
		last = elem
		found = true
	})

	return last, found
}

func (t *traversable[T]) LastOption() *Option[T] {
	last, found := t.Last()

	if found {
		return Some(last)
	}

	return None[T]()
}

func (t *traversable[T]) Find(filter func(elem T) bool) *Option[T] {
	var found bool
	var result *Option[T] = None[T]()

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

func (t *traversable[T]) Exists(filter func(elem T) bool) bool {
	return t.Find(filter).Count(nil) > 0
}

func (t *traversable[T]) Count(filter func(elem T) bool) int {
	var result int
	t.ForEach(func(elem T) {
		if filter == nil || filter(elem) {
			result++
		}
	})

	return result
}

func (t *traversable[T]) ForAll(filter func(elem T) bool, transform func(elem T) T) *List[T] {
	newList := EmptyList[T]()

	t.ForEach(func(elem T) {
		if filter == nil || filter(elem) {
			newList.Add(transform(elem))
		} else {
			newList.Add(elem)
		}
	})

	return newList
}

func (t *traversable[T]) IsEmpty() bool {
	return t.HeadOption().Count(nil) == 0
}

func (t *traversable[T]) NonEmpty() bool {
	return !t.IsEmpty()
}

func (t *traversable[T]) IsDefined() bool {
	return !t.IsEmpty()
}

func (t *traversable[T]) Reduce(op func(a, b T) T) T {
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

func (t *traversable[T]) Select(filter func(elem T) bool) *List[T] {
	newList := EmptyList[T]()

	t.ForEach(func(elem T) {
		if filter(elem) {
			newList.Add(elem)
		}
	})

	return newList
}

func (t *traversable[T]) ToList() *List[T] {
	newList := EmptyList[T]()

	t.ForEach(func(elem T) {
		newList.Add(elem)
	})

	return newList
}

// Special cases which can't be implemented using the current golang

func ToSet[T comparable](t *traversable[T]) *Set[T] {
	newSet := EmptySet[T]()
	t.ForEach(func(elem T) {
		newSet.Add(elem)
	})

	return newSet
}

func Map[T, V any](t Traversable[T], transform func(T) V) Traversable[V] {
	newList := EmptyList[V]()
	t.ForEach(func(elem T) {
		newList.Add(transform(elem))
	})

	return newList
}

type number interface {
	constraints.Integer | constraints.Float
}

func Sum[T number](t *traversable[T]) T {
	var result T
	t.ForEach(func(elem T) {
		result += elem
	})

	return result
}

func Product[T number](t *traversable[T]) T {
	var result T
	t.ForEach(func(elem T) {
		result *= elem
	})

	return result
}

func Min[T constraints.Ordered](t *traversable[T]) T {
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

func Max[T constraints.Ordered](t *traversable[T]) T {
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

func Fold[T, V any](t Traversable[T], initialValue V, op func(a V, b T) V) V {
	accumulator := initialValue

	t.ForEach(func(elem T) {
		accumulator = op(accumulator, elem)
	})

	return accumulator
}
