package collection

type Option[T any] struct {
	traversable[T]
	value *T
}

var _ Traversable[string] = &Option[string]{}

func NewOption[T any](value *T) *Option[T] {
	opt := &Option[T]{value: value}
	opt.forEach = opt
	opt.traversable.forEach = opt

	return opt
}

func Some[T any](value T) *Option[T] {
	return NewOption[T](&value)
}

func None[T any]() *Option[T] {
	return NewOption[T](nil)
}

func (o *Option[T]) ForEach(action func(T)) {
	if o.value != nil {
		action(*o.value)
	}
}

func (o *Option[T]) GetOrElse(altValue T) T {
	head, found := o.Head()
	if found {
		return head
	}

	return altValue
}

func Ptr[T any](t T) *T {
	return &t
}
