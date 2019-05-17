package main

import "golang.org/x/tour/tree"

import "fmt"

// WalkRer recursive function
func WalkRer(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkRer(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkRer(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRer(t, ch)
	close(ch)
}

func readTree(ch chan int) (result []int) {
	for val := range ch {
		result = append(result, val)
	}
	return result
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	vals1 := readTree(ch1)
	vals2 := readTree(ch2)
	if len(vals1) != len(vals2) {
		return false
	}
	for idx, val := range vals1 {
		if val != vals2[idx] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
