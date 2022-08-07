package collection

type KV[K comparable, V any] map[K]V

func (l KV[K, V]) ForEach(action func(tuple Tuple2[K, V])) {
	for key, value := range l {
		action(Tuple2[K, V]{Field1: key, Field2: value})
	}
}

func NewKV[K comparable, V any](elements ...Tuple2[K, V]) KV[K, V] {
	if len(elements) == 0 {
		return make(KV[K, V])
	}

	newKV := make(KV[K, V], len(elements))
	for _, elem := range elements {
		newKV[elem.Field1] = elem.Field2
	}

	return newKV
}

func EmptyKV[K comparable, V any]() KV[K, V] {
	return NewKV[K, V]()
}
