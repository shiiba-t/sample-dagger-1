package main

import "fmt"

func main() {
	fmt.Printf("Sum: 1 + 2 = %d\n", sum(1, 2))
}

func sum(a, b int64) int64 {
	return a + b
}
