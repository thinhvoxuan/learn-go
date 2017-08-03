package chainReaction1

import (
	"reflect"
	"testing"
)

const targetTestVersion = 2

type testCaseCreateEmptyArray struct {
	input    int
	expected [][]int
}

var testcase1 = []testCaseCreateEmptyArray{
	{1, [][]int{{0}}},
	{3, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
}

func TestCreateEmptyArray(t *testing.T) {
	for _, test := range testcase1 {
		actual := createEmptyArray(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("CreateEmpty test [%d], expected [%v], actual [%v]", test.input, test.expected, actual)
		}
	}
}

type testCaseMoveStep struct {
	currentMap [][]int
	move       []int
	expected   [][]int
}

var testcase2 = []testCaseMoveStep{
	{[][]int{{0, 0}, {0, 0}}, []int{0, 0}, [][]int{{1, 0}, {0, 0}}},
	{[][]int{{0, 0}, {0, 0}}, []int{1, 1}, [][]int{{0, 0}, {0, 1}}},
	{[][]int{{1, 0}, {0, 0}}, []int{0, 0}, [][]int{{0, 1}, {1, 0}}},
}

func TestAddPoint(t *testing.T) {
	for _, test := range testcase2 {
		actual := addPoint(test.currentMap, test.move, 2)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("AddPoint test [%d], expected [%v], actual [%v]", test.currentMap, test.expected, actual)
		}
	}
}

type testGetMaxPosition struct {
	dx       int
	dy       int
	n        int
	expected int
}

var testcase3 = []testGetMaxPosition{
	{0, 0, 4, 2},
	{3, 3, 4, 2},
	{0, 3, 4, 2},
	{3, 0, 4, 2},
	{2, 2, 4, 4},
	{2, 3, 4, 3},
}

func TestGetMaxDimenssionInPosition(t *testing.T) {
	for _, test := range testcase3 {
		actual := getMaxDimenssionInPosition(test.dx, test.dy, test.n)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("getMaxDimenssionInPosition test [%d] [%d] [%d], expected [%v], actual [%v]", test.dx, test.dy, test.n, test.expected, actual)
		}
	}
}
