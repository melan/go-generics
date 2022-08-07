package collection

type Tuple1[T1 any] struct {
	Field1 T1
}

type Tuple2[T1, T2 any] struct {
	Field1 T1
	Field2 T2
}

type Tuple3[T1, T2, T3 any] struct {
	Field1 T1
	Field2 T2
	Field3 T3
}

type Tuple4[T1, T2, T3, T4 any] struct {
	Field1 T1
	Field2 T2
	Field3 T3
	Field4 T4
}
