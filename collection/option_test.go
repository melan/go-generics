package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrElse(t *testing.T) {
	tests := []struct {
		name     string
		value    *Option[string]
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
			got := tt.value.GetOrElse(tt.altValue)
			assert.Equal(t, tt.want, got)
		})
	}
}
