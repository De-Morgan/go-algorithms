package arrays

import "testing"

func Test_lowestValue(t *testing.T) {
	type args struct {
		value []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name:  "empty_values",
			args:  args{},
			want:  0,
			want1: false,
		},
		{
			name: "lowest value at first index",
			args: args{
				value: []int{
					1, 3, 5, 2, 6, 8, 4,
				},
			},
			want:  1,
			want1: true,
		},
		{
			name: "lowest value at last index",
			args: args{
				value: []int{
					1, 3, 5, 2, 6, -7, 8, 4, -50,
				},
			},
			want:  -50,
			want1: true,
		},
		{
			name: "lowest value in the middle",
			args: args{
				value: []int{
					1, 3, 5, 2, 6, 0, 8, 4, 9,
				},
			},
			want:  0,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := lowestValue(tt.args.value)
			if got != tt.want {
				t.Errorf("lowestValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("lowestValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
