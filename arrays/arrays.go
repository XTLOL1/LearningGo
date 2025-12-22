package main

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

// acts as two dimentional array
func sumAllNumbers[T Number](arraysToSum ...[]T) (sum []T) {
	for _, numbers := range arraysToSum {
		sum = append(sum, sumNumbers(numbers))
	}
	return
}

func sumAllNumbersTails[T Number](arraysToSum ...[]T) (sum []T) {
	for _, numbers := range arraysToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, sumNumbers(tail))
		}
	}
	return
}
