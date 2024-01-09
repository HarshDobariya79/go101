package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println("The value of ptr is", ptr)

	a := 5

	ptr = &a

	fmt.Println("The value of ptr is", ptr)
	fmt.Println("ptr points to the value", *ptr)
}
