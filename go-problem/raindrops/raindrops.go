package raindrops

import (
	"strconv"
)

const testVersion = 2

//Convert return int to string
func Convert(input int) string {
	result := ""
	if input%3 == 0 {
		result = result + "Pling"
	}
	if input%5 == 0 {
		result = result + "Plang"
	}
	if input%7 == 0 {
		result = result + "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(input)
	}
	return result
}
