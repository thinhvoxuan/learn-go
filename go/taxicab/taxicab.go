package taxicab

import (
	"strconv"
	"strings"
)

// Move struct
type Move struct {
	direction string
	step      int
}

// Vector2D struct
type Vector2D struct {
	x int
	y int
}

func (vector Vector2D) add(x int) Vector2D {
	vector.x = vector.x + x
	return vector
}

func (vector Vector2D) rotate(direction string) Vector2D {
	var xValue = vector.x
	var yValue = vector.y
	if direction == "R" {
		vector.x = (-1) * yValue
		vector.y = xValue
	} else {
		vector.x = yValue
		vector.y = (-1) * xValue
	}
	return vector
}

// MappingInputToArrayMove mapping
func MappingInputToArrayMove(input string) []Move {
	var moveArray = make([]Move, 0)
	var arrayString = strings.Split(input, ", ")
	for _, str := range arrayString {
		var direction = string(str[0])
		move, _ := strconv.Atoi(str[1:])
		moveArray = append(moveArray, Move{direction, move})
	}
	return moveArray
}

func absInt(x int) int {
	if x > 0 {
		return x
	}
	return (-1) * x
}

// BlockAway calulating the nearest block from input string
func BlockAway(input string) int {
	var moveArray = MappingInputToArrayMove(input)
	var vectorMove = Vector2D{0, 0}
	for _, move := range moveArray {
		vectorMove = vectorMove.rotate(move.direction)
		vectorMove = vectorMove.add(move.step)
	}
	return absInt(vectorMove.x) + absInt(vectorMove.y)
}
