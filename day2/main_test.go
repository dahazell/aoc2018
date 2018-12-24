package main

import (
	"testing"
)

func Test_findRepeats(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantDoubles int
		wantTriples int
	}{
		{
			name: "no matches",
			args: args{
				s: "abcdef",
			},
			wantDoubles: 0,
			wantTriples: 0,
		},
		{
			name: "find double",
			args: args{
				s: "abbcde",
			},
			wantDoubles: 1,
			wantTriples: 0,
		},
		{
			name: "find both",
			args: args{
				s: "bababc",
			},
			wantDoubles: 1,
			wantTriples: 1,
		},
		{
			name: "two triples don't count doubles",
			args: args{
				s: "ababab",
			},
			wantDoubles: 0,
			wantTriples: 1,
		},
		{
			name: "four the same is 0 for all",
			args: args{
				s: "bbbb",
			},
			wantDoubles: 0,
			wantTriples: 0,
		},
		{
			name: "one double ignore 2 sets of four",
			args: args{
				s: "bbbbdeeffee",
			},
			wantDoubles: 1,
			wantTriples: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDoubles, gotTriples := findRepeats(tt.args.s)
			if gotDoubles != tt.wantDoubles {
				t.Errorf("findRepeats() gotDoubles = %v, want %v", gotDoubles, tt.wantDoubles)
			}
			if gotTriples != tt.wantTriples {
				t.Errorf("findRepeats() gotTriples = %v, want %v", gotTriples, tt.wantTriples)
			}
		})
	}
}

func Test_calculateCheckSum(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "work out total",
			args: args{
				input: []string{
					"abcdef",
					"bababc",
					"abbcde",
					"abcccd",
					"aabcdd",
					"abcdee",
					"ababab",
				},
			},
			want: 12,
		},
		{
			name: "test",
			args: args{
				input: []string{
					"abcdefghijkmnoooqrrstvwxyz",
					"abcefghiijklmnopqrstuvwxxy",
					"abccccefghijklmnoqrstuvwxy",
					"abdefgijjklmnoprstuvvwwxyz",
					"bccdfghhijklmnopqrstuvwxyz",
					"abceghijklmnopqqrstuwxyyyz",
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCheckSum(tt.args.input); got != tt.want {
				t.Errorf("calculateCheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
