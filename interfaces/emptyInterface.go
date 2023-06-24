package main

import (
	"fmt"
)

func JustPrint(i interface{}) {
	fmt.Printf("Type: %T value: %v", i, i)
}

func main() {
	JustPrint("hello");
	JustPrint(10);
	JustPrint(12.99);
	// JustPrint({1,2,3})

}
