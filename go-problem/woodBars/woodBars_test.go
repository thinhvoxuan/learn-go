package main

import (
	"reflect"
	"testing"
)

func Test_woodBars(t *testing.T) {
	type args struct {
		bars []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{4, 5, 6}}, 1},
		{"t2", args{[]int{10}}, 10},
		{"t3", args{[]int{6, 16, 12, 2, 18}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := woodBars(tt.args.bars); got != tt.want {
				t.Errorf("woodBars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reduceValue(t *testing.T) {
	type args struct {
		bars []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"t1", args{[]int{4, 5, 6}}, []int{4, 1, 2}},
		{"t1", args{[]int{2, 6, 16, 12, 18}}, []int{2, 4, 16, 10, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := reduceValue(tt.args.bars); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("reduceValue() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
