package main

import (
	"testing"
)

func Test_calculateFrequencyTotal(t *testing.T) {
	type args struct {
		frequencies []string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int64
		wantErr   bool
	}{
		{
			name: "Gold path all additions",
			args: args{
				frequencies: []string{
					"+1", "+1", "+1",
				},
			},
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "Gold path mixed",
			args: args{
				frequencies: []string{
					"+1", "+1", "-2",
				},
			},
			wantTotal: 0,
			wantErr:   false,
		},
		{
			name: "Gold path all subtractions",
			args: args{
				frequencies: []string{
					"-1", "-1", "-2",
				},
			},
			wantTotal: -4,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, err := calculateFrequencyTotal(tt.args.frequencies)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateFrequencyTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("calculateFrequencyTotal() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_firstFrequencyFoundTwice(t *testing.T) {
	type args struct {
		frequencies []string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int64
		wantErr   bool
	}{
		{
			name: "Zero found twice",
			args: args{
				frequencies: []string{
					"+1", "-1",
				},
			},
			wantTotal: 0,
			wantErr:   false,
		}, //+3, +3, +4, -2, -4
		{
			name: "Ten found twice",
			args: args{
				frequencies: []string{
					"+3", "+3", "+4", "-2", "-4",
				},
			},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, err := firstFrequencyFoundTwice(tt.args.frequencies)
			if (err != nil) != tt.wantErr {
				t.Errorf("firstFrequencyFoundTwice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("firstFrequencyFoundTwice() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
