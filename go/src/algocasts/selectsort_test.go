package algocasts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectSortMin(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"single", []int{1, 5, 4}, []int{1, 4, 5}},
		{"normal", []int{1, 9, 5, 3, 2, 6}, []int{1, 2, 3, 5, 6, 9}},
	}
	for _, tt := range tests {
		selectsortMin(tt.input)
		assert.Equal(t, tt.input, tt.want)
	}
}

func TestSelectSortMax(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"single", []int{1, 5, 4}, []int{1, 4, 5}},
		{"normal", []int{1, 9, 5, 3, 2, 6}, []int{1, 2, 3, 5, 6, 9}},
	}
	for _, tt := range tests {
		selectsortMax(tt.input)
		assert.Equal(t, tt.input, tt.want)
	}
}
