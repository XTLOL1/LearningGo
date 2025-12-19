package main

import "fmt"

func iterateString(iterable string) (returnValue string) {
	for i := 0; i<5; i++ {
		returnValue += iterable
	}
	return
}

func main() {
	fmt.Println(iterateString("sneeze"))
}
