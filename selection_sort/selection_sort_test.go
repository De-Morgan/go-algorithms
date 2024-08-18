package main

import (
	"slices"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty_slice",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "already_sorted",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "not_sorted",
			args: args{
				nums: []int{3, 1, 2, 5, 4},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "repeated_numbers",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.nums
			selectionSort(got)
			if !slices.Equal(tt.want, got) {
				t.Errorf("selectionSort() got = %v, want %v", got, tt.want)
			}
		})
	}
}
