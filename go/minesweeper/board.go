package minesweeper

import "bytes"

const testVersion = 1

//Board object
type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

//Count the board
func (b Board) Count() (err error) {
	for x, r := range b {
		for y := range r {
			b[x][y] = '1'
		}
	}
	return nil
}
