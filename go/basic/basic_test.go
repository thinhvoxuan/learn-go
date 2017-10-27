package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{"testcase 1", args{1, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := add(tt.args.a, tt.args.b); gotC != tt.wantC {
				t.Errorf("add() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
