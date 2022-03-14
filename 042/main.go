// ● Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned func
package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
