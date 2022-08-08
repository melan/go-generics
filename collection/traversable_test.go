package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFold(t *testing.T) {
	type args struct {
		t            Traversable[string]
		initialValue int
		op           func(a int, b string) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count total length of the list",
			args: args{
				t:            NewList[string]("hi", "world"),
				initialValue: 0,
				op: func(a int, b string) int {
					return a + len(b)
				},
			},
			want: 7,
		},
		{
			name: "count total length of defined optional",
			args: args{
				t:            Some[string]("hello"),
				initialValue: 0,
				op: func(a int, b string) int {
					return a + len(b)
				},
			},
			want: 5,
		},
		{
			name: "count total length of an empty optional",
			args: args{
				t:            None[string](),
				initialValue: 0,
				op: func(a int, b string) int {
					return a + len(b)
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Fold(tt.args.t, tt.args.initialValue, tt.args.op), "Fold(%v, %v, %v)", tt.args.t, tt.args.initialValue, tt.args.op)
		})
	}
}

func TestChain(t *testing.T) {
	pipeline := func(t Traversable[string]) int {
		optionals := Map[string, *Option[string]](t, func(elem string) *Option[string] {
			return Some[string](elem)
		})

		return optionals.
			ForAll(
				func(elem *Option[string]) bool {
					return elem.IsEmpty() || elem.GetOrElse("") != "hello"
				},
				func(elem *Option[string]) *Option[string] {
					return None[string]()
				}).
			Select(func(elem *Option[string]) bool {
				return elem.IsDefined()
			}).
			Count(nil)
	}

	tests := []struct {
		name    string
		initial Traversable[string]
		want    int
	}{
		{
			name:    "defined optional",
			initial: Some[string]("hello"),
			want:    1,
		},
		{
			name:    "alt defined optional",
			initial: Some[string]("world"),
			want:    0,
		},
		{
			name:    "list with the value",
			initial: NewList("world", "hello", "hi", "hello", "again"),
			want:    2,
		},
		{
			name:    "list without the value",
			initial: NewList("world", "hi"),
			want:    0,
		},
		{
			name:    "set with the value",
			initial: NewSet("world", "hello", "hi", "hello", "again"),
			want:    1,
		},
		{
			name:    "set without the value",
			initial: NewSet("world", "hi"),
			want:    0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := pipeline(tt.initial)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	transformation := func(s string) int {
		return len(s)
	}

	tests := []struct {
		name     string
		value    Traversable[string]
		want     int
		validate func(t *testing.T, result Traversable[int])
	}{
		{
			name:  "optional, some text",
			value: Some[string]("hello"),
			want:  5,
		},
		{
			name:  "optional, none",
			value: NewOption[string](nil),
			want:  0,
		},
		{
			name:  "empty list",
			value: EmptyList[string](),
			want:  0,
		},
		{
			name:  "non-empty list",
			value: NewList("foo", "bar"),
			want:  6,
		},
		{
			name:  "non-empty list with duplicates",
			value: NewList("foo", "bar", "foo"),
			want:  9,
		},
		{
			name:  "empty set",
			value: EmptySet[string](),
			want:  0,
		},
		{
			name:  "non-empty set",
			value: NewSet("foo", "bar"),
			want:  6,
		},
		{
			name:  "non-empty set with duplicates",
			value: NewSet("foo", "bar", "bar"),
			want:  6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := Map[string, int](tt.value, transformation).Reduce(func(a, b int) int { return a + b })
			assert.Equal(t, tt.want, got)
		})
	}
}
