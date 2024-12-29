package main

import "fmt"

func main() {
	fmt.Println("Test.")
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
