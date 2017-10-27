package taxicab

import (
	"reflect"
	"testing"
)

func TestBlockAway(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"testcase1", args{"R2, L3"}, 5},
		{"testcase2", args{"R2, R2, R2"}, 2},
		{"testcase3", args{"R5, L5, R5, R3"}, 12},
		{"testcase4", args{"R5, R4, R2, L3, R1, R1, L4, L5, R3, L1, L1, R4, L2, R1, R4, R4, L2, L2, R4, L4, R1, R3, L3, L1, L2, R1, R5, L5, L1, L1, R3, R5, L1, R4, L5, R5, R1, L185, R4, L1, R51, R3, L2, R78, R1, L4, R188, R1, L5, R5, R2, R3, L5, R3, R4, L1, R2, R2, L4, L4, L5, R5, R4, L4, R2, L5, R2, L1, L4, R4, L4, R2, L3, L4, R2, L3, R3, R2, L2, L3, R4, R3, R1, L4, L2, L5, R4, R4, L1, R1, L5, L1, R3, R1, L2, R1, R1, R3, L4, L1, L3, R2, R4, R2, L2, R1, L5, R3, L3, R3, L1, R4, L3, L3, R4, L2, L1, L3, R2, R3, L2, L1, R4, L3, L5, L2, L4, R1, L4, L4, R3, R5, L4, L1, L1, R4, L2, R5, R1, R1, R2, R1, R5, L1, L3, L5, R2"}, 231},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlockAway(tt.args.input); got != tt.want {
				t.Errorf("BlockAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMappingInputToArrayMove(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []Move
	}{
		{
			"test1", args{"R2, L3"}, []Move{
				Move{"R", 2},
				Move{"L", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MappingInputToArrayMove(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MappingInputToArrayMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
