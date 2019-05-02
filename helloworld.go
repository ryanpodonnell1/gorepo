package main

import (
	"fmt"
)

var y = 42
var z = "Shaken, not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
