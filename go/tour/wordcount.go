package main

import "strings"

// WordCount return workcount in string
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	ss := strings.Split(s, " ")
	for _, value := range ss {
		if val, ok := result[value]; ok {
			result[value] = val + 1
		} else {
			result[value] = 1
		}
	}
	return result
}

// func main() {
// 	fmt.Println(WordCount("I am learning Go!"))
// }
