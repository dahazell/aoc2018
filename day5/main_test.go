package main

import (
	"testing"
)

func Test_reactPolymers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple react",
			args: args{
				input: "dabAcCaCBAcCcaDA",
			},
			want: "dabCBAcaDA",
		},
		{
			name: "first chars react",
			args: args{
				input: "dDabAcCaCBAcCcaDA",
			},
			want: "abCBAcaDA",
		},
		{
			name: "multiple same char",
			args: args{
				input: "dDabAcCcCcCcCaCBAcCcaDA ",
			},
			want: "abCBAcaDA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reactPolymers(tt.args.input); got != tt.want {
				t.Errorf("reactPolymers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestPolymerRemovedUnits(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name        string
		args        args
		wantLen     int
		wantRemoved string
	}{
		{
			name: "simple react",
			args: args{
				input: "dabAcCaCBAcCcaDA",
			},
			wantLen:     4,
			wantRemoved: "Cc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := shortestPolymerRemovedUnits(tt.args.input)
			if got != tt.wantLen {
				t.Errorf("shortestPolymerRemovedUnits() got = %v, want %v", got, tt.wantLen)
			}
			if got1 != tt.wantRemoved {
				t.Errorf("shortestPolymerRemovedUnits() got1 = %v, want %v", got1, tt.wantRemoved)
			}
		})
	}
}
