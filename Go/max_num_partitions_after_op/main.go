package main

import "fmt"

func maxPartitionsAfterOperations(s string, k int) int {
	if k > len(s) {
		return 1
	}

}

func main() {
	s := "accca"
	k := 2
	result := maxPartitionsAfterOperations(s, k)
	fmt.Println(result)
}
