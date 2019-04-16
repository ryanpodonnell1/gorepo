package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	foo()
	fmt.Println("Turd Turd Turd")

	for i := 0; i < 100; i++ {
		if i%9 == 0 && i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func foo() {

	fmt.Println("I'm in foo!")

}

//Control flow:
// (1) sequence
// (2) Loop
// (3) Conditional
