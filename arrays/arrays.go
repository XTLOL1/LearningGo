package main

import (
	"fmt"
)

// Define a constraint for all types that support arithmetic
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func sumNumbers[T Number](arrayNumbers []T) (sum T) {
	for _, number := range arrayNumbers {
		sum += number
	}
	return
}

func main() {
	fmt.Println(sumNumbers([]int{1, 2, 3}))
}
