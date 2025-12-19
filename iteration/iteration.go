package main

import (
	"fmt"
	"strings"
)

func iterateString(iterable string, repeatCount int) string {
	var returnValue strings.Builder
	for i := 0; i < repeatCount; i++ {
		returnValue.WriteString(iterable)
	}
	return returnValue.String()
}

func main() {
	fmt.Println(iterateString("sneeze", 0))
}
