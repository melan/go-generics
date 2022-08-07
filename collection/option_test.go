package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name           string
		value          Option[string]
		transformation func(string) int
		validate       func(t *testing.T, result Traversable[int])
	}{
		{
			name:  "some text",
			value: Some[string]("hello"),
			transformation: func(s string) int {
				return len(s)
			},
			validate: func(t *testing.T, result Traversable[int]) {
				assert.Equal(t, 5, GetOrElse[int](result, -1))
			},
		},
		{
			name:  "none",
			value: NewOption[string](nil),
			transformation: func(s string) int {
				return len(s)
			},
			validate: func(t *testing.T, result Traversable[int]) {
				assert.Equal(t, -1, GetOrElse[int](result, -1))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := Map[string, int](tt.value, tt.transformation)
			tt.validate(t, got)
		})
	}
}

func TestGetOrElse(t *testing.T) {
	tests := []struct {
		name     string
		value    Option[string]
		altValue string
		want     string
	}{
		{
			name:     "some text",
			value:    Some[string]("hello"),
			altValue: "test test",
			want:     "hello",
		},
		{
			name:     "none",
			value:    NewOption[string](nil),
			altValue: "test test",
			want:     "test test",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := GetOrElse[string](tt.value, tt.altValue)
			assert.Equal(t, tt.want, got)
		})
	}
}
