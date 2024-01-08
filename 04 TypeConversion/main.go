package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the amount: ")

	input, _ := reader.ReadString('\n')

	fmt.Printf("The type of the amount is: %T\n", input)

	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic(err)
	}

	amount = amount + amount*0.18

	fmt.Println("The amount with GST is:", amount)
}
