package main

import (
	"reflect"
	"testing"
)

func Test_findSleepyGuard(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want guard
	}{
		{
			name: "small sample input",
			args: args{
				input: []string{
					"[1518-11-01 00:00] Guard #10 begins shift",
					"[1518-11-01 00:05] falls asleep",
					"[1518-11-01 00:25] wakes up",
					"[1518-11-01 00:30] falls asleep",
					"[1518-11-01 00:55] wakes up",
					"[1518-11-01 23:58] Guard #99 begins shift",
					"[1518-11-02 00:40] falls asleep",
					"[1518-11-02 00:50] wakes up",
					"[1518-11-03 00:05] Guard #10 begins shift",
					"[1518-11-03 00:24] falls asleep",
					"[1518-11-03 00:29] wakes up",
					"[1518-11-04 00:02] Guard #99 begins shift",
					"[1518-11-04 00:36] falls asleep",
					"[1518-11-04 00:46] wakes up",
					"[1518-11-05 00:03] Guard #99 begins shift",
					"[1518-11-05 00:45] falls asleep",
					"[1518-11-05 00:55] wakes up",
				},
			},
			want: guard{
				id:         10,
				minute:     24,
				timeAsleep: 50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSleepyGuard(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSleepyGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMostFrequentMinAsleep(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want guard
	}{
		{
			name: "Find most freq min",
			args: args{
				input: []string{
					"[1518-11-01 00:00] Guard #10 begins shift",
					"[1518-11-01 00:05] falls asleep",
					"[1518-11-01 00:25] wakes up",
					"[1518-11-01 00:30] falls asleep",
					"[1518-11-01 00:55] wakes up",
					"[1518-11-01 23:58] Guard #99 begins shift",
					"[1518-11-02 00:40] falls asleep",
					"[1518-11-02 00:50] wakes up",
					"[1518-11-03 00:05] Guard #10 begins shift",
					"[1518-11-03 00:24] falls asleep",
					"[1518-11-03 00:29] wakes up",
					"[1518-11-04 00:02] Guard #99 begins shift",
					"[1518-11-04 00:36] falls asleep",
					"[1518-11-04 00:46] wakes up",
					"[1518-11-05 00:03] Guard #99 begins shift",
					"[1518-11-05 00:45] falls asleep",
					"[1518-11-05 00:55] wakes up",
				},
			},
			want: guard{
				id:         99,
				minute:     45,
				minuteFreq: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMostFrequentMinAsleep(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMostFrequentMinAsleep() = %v, want %v", got, tt.want)
			}
		})
	}
}
