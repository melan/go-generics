package collection

type Option[T any] struct {
	value *T
}

func NewOption[T any](value *T) Option[T] {
	return Option[T]{value: value}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return NewOption[T](nil)
}

func (o Option[T]) ForEach(action func(T)) {
	if o.value != nil {
		action(*o.value)
	}
}

func GetOrElse[T any](opt Traversable[T], altValue T) T {
	head, found := Head[T](opt)
	if found {
		return head
	}

	return altValue
}

func Ptr[T any](t T) *T {
	return &t
}
