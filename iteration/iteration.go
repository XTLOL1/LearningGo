package main

import "fmt"

func iterateString(iterable string, repeatCount uint) string {
	var returnValue strings.Builder
	for i := 0; i < repeatCount; i++ {
		returnValue.WriteString(iterable)
	}
	return returnValue.String()
}

func main() {
	fmt.Println(iterateString("sneeze"))
}
