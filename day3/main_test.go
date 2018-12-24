package main

import "testing"

func Test_findOverlaps(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test 3 inputs",
			args: args{
				lines: []string{
					"#1 @ 1,3: 4x4",
					"#2 @ 3,1: 4x4",
					"#3 @ 5,5: 2x2",
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOverlaps(tt.args.lines); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
