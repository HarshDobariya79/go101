package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your name: ")

	name, _ := reader.ReadString('\n')

	fmt.Println("Hello", name)

	// Numbers entered as input are considered as a string
}
