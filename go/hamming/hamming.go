package hamming

import (
	"errors"
)

const testVersion = 5

// Distance between 2 DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Could not compare")
	}
	total := 0
	for idx, chr := range a {
		if string(chr) != string(b[idx]) {
			total = total + 1
		}
	}
	return total, nil
}
