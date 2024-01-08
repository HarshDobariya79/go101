package main

import "fmt"

const LoginURL = "https://github.com/harshdobariya79"
// Global variables names should start with a Capital letter

func main() {
	var name string = "Harsh" // we can also use var name = "Harsh" or name := "Harsh"
	// := operator declarations are only allowed inside any method

	fmt.Printf("name is of type %T\n", name)
	fmt.Println("Hello from", name)

	var isSignedIn bool = false
	fmt.Println("Signed in?", isSignedIn)

	var unassignedVar int
	fmt.Println("The value of unassignedVar is:", unassignedVar)
}
