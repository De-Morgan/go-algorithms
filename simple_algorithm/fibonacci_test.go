package main

import (
	"reflect"
	"testing"
)

func Test_fibonacciLoop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "first five fibonacci series",
			args: args{
				n: 5,
			},
			want: []uint64{
				0, 1, 1, 2, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciLoop(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacciLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacciRecursion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "first five fibonacci series",
			args: args{
				n: 5,
			},
			want: []uint64{
				0, 1, 1, 2, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciRecursion(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacciRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacciN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "10th fibonacci",
			args: args{
				n: 10,
			},
			want: 55,
		},
		{
			name: "18th fibonacci",
			args: args{
				n: 18,
			},
			want: 2584,
		},
		{
			name: "28th fibonacci",
			args: args{
				n: 28,
			},
			want: 317811,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciN(tt.args.n); got != tt.want {
				t.Errorf("fibonacciN() = %v, want %v", got, tt.want)
			}
		})
	}
}
