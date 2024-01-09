package main

import "fmt"

func main() {
	var nums [3]int

	fmt.Println(nums)

	nums[0] = 5
	nums[2] = 2

	fmt.Println(nums)
	fmt.Println(len(nums))

	var alphabets = [3]string{"a", "b", "c"}

	fmt.Println(alphabets)
}
