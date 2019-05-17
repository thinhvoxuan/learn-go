package main

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev2, prev1 := 0, 1
	return func() int {
		prev2, prev1 = prev1+prev2, prev2
		return prev2
	}
}

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
