package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name:     "",
			args:     args{arr: []int{5, 2, 4, 6, 1, 3}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, InsertSort(tt.args.arr))
		})
	}
}
